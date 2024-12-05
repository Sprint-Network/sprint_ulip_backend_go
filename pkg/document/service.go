package document

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"prechecks/pkg"
)

func FetchEmailFromCIN(cin string) (string, error) {
	apiURL := "https://www.ulipstaging.dpiit.gov.in/ulip/v1.0.0/MCA/03"

	payload := map[string]string{"CIN": cin}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to serialize request payload: %w", err)
	}

	client := &http.Client{}

	// Helper function to make the request with a given token
	makeRequest := func(token string) (*http.Response, error) {
		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		return client.Do(req)
	}

	// Try the first request with the current token
	token, err := pkg.GetAuthToken()
	if err != nil {
		return "", fmt.Errorf("failed to get auth token: %w", err)
	}

	resp, err := makeRequest(token)
	if err != nil {
		return "", fmt.Errorf("request to MCA API failed: %w", err)
	}
	defer resp.Body.Close()

	// If the response is 403, fetch a new token and retry
	if resp.StatusCode == http.StatusForbidden {
		token, err = pkg.FetchAuthToken()
		if err != nil {
			return "", fmt.Errorf("failed to refresh auth token: %w", err)
		}

		resp, err = makeRequest(token)
		if err != nil {
			return "", fmt.Errorf("retry request to MCA API failed: %w", err)
		}
		defer resp.Body.Close()
	}

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("MCA API returned status %d: %s", resp.StatusCode, string(body))
	}

	var apiResponse struct {
		Response []struct {
			Response struct {
				Data []struct {
					EmailAddress string `json:"emailAddress"`
				} `json:"data"`
			} `json:"response"`
		} `json:"response"`
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return "", fmt.Errorf("failed to parse response JSON: %w", err)
	}

	// Extract email address
	if len(apiResponse.Response) > 0 &&
		len(apiResponse.Response[0].Response.Data) > 0 &&
		apiResponse.Response[0].Response.Data[0].EmailAddress != "" {
		return apiResponse.Response[0].Response.Data[0].EmailAddress, nil
	}

	return "", fmt.Errorf("email address not found in response")
}
