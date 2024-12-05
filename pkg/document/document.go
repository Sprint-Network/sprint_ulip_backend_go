package document

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// Regex to match CIN numbers
var cinRegex = regexp.MustCompile(`\b[UL][0-9]{5}[A-Za-z]{2}[0-9]{4}[A-Za-z]{3}\s?[0-9]{6}\b`)

// Extract CIN from the provided PDF file using pdfcpu, with OCR fallback for image-based PDFs
func ExtractCINFromPDF(file io.Reader) (string, error) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %v", err)
	}

	// Create a temporary file to store the uploaded PDF
	tmpFile, err := os.CreateTemp(cwd, "uploaded-*.pdf")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %v", err)
	}
	// Defer the cleanup to remove the file once the function exits
	defer func() {
		if tmpFile != nil {
			err := os.Remove(tmpFile.Name())
			if err != nil {
				log.Printf("failed to remove temp file: %v", err)
			}
		}
	}()

	// Ensure tmpFile gets properly closed
	defer tmpFile.Close()

	// Copy the uploaded content to the temporary file
	if _, err := io.Copy(tmpFile, file); err != nil {
		return "", fmt.Errorf("failed to save uploaded file: %v", err)
	}

	// Create a temporary directory to store the extracted content
	tmpDir, err := os.MkdirTemp(cwd, "extracted-content-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up extracted content

	// Try extracting content using pdfcpu
	err = api.ExtractContentFile(tmpFile.Name(), tmpDir, nil, nil)
	if err != nil {
		// Fallback to OCR if pdfcpu fails to extract content
		ocrText, ocrErr := performOCR(tmpFile.Name())
		if ocrErr != nil {
			return "", fmt.Errorf("failed to extract content from PDF: %v, and OCR failed: %v", err, ocrErr)
		}

		// Search for the CIN in OCR-extracted text
		matches := cinRegex.FindStringSubmatch(ocrText)
		if len(matches) > 0 {
			// Remove any spaces in the matched CIN (if present)
			cin := strings.ReplaceAll(matches[0], " ", "")
			return cin, nil
		}
		return "", fmt.Errorf("CIN not found in OCR-processed document")
	}

	// Read the extracted text file
	textFile := filepath.Join(tmpDir, "content.txt")
	if _, err := os.Stat(textFile); os.IsNotExist(err) {
		// Fallback to OCR if content.txt is missing
		ocrText, ocrErr := performOCR(tmpFile.Name())
		if ocrErr != nil {
			return "", fmt.Errorf("content.txt not found, and OCR failed: %v", ocrErr)
		}

		// Search for the CIN in OCR-extracted text
		matches := cinRegex.FindStringSubmatch(ocrText)
		if len(matches) > 0 {
			// Remove any spaces in the matched CIN (if present)
			cin := strings.ReplaceAll(matches[0], " ", "")
			return cin, nil
		}
		return "", fmt.Errorf("CIN not found in OCR-processed document")
	}

	content, err := os.ReadFile(textFile)
	if err != nil {
		return "", fmt.Errorf("failed to read extracted text: %v", err)
	}

	// Search for the CIN in the extracted text
	matches := cinRegex.FindStringSubmatch(string(content))
	if len(matches) > 0 {
		return matches[0], nil
	}

	return "", fmt.Errorf("CIN not found in the document")
}

func performOCR(pdfPath string) (string, error) {
	// Convert the PDF to images using pdftoppm (requires poppler-utils)
	imageBase := pdfPath + "-page"
	cmd := exec.Command("pdftoppm", "-jpeg", pdfPath, imageBase)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to convert PDF to images: %v", err)
	}

	// Perform OCR on each image
	var extractedText string
	pagePattern := filepath.Join(filepath.Dir(pdfPath), filepath.Base(imageBase)+"*.jpg")
	matches, err := filepath.Glob(pagePattern)
	if err != nil || len(matches) == 0 {
		return "", fmt.Errorf("no images generated from PDF: %v", err)
	}

	// Perform OCR on all matched image files
	for _, img := range matches {
		cmd := exec.Command("tesseract", img, "stdout")
		output, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("failed to perform OCR on image %s: %v", img, err)
		}
		extractedText += string(output)

		// Delete the image file after OCR processing
		if err := os.Remove(img); err != nil {
			log.Printf("failed to remove image file %s: %v", img, err)
		}
	}

	// Return the extracted text after processing all images
	return extractedText, nil
}
