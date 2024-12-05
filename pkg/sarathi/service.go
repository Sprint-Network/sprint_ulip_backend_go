package sarathi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"prechecks/pkg"
	"prechecks/pkg/sarathi/models"
	"strings"
)

func ValidateDLWithUIP(dlNumber, dob, name string) (bool, error) {
	url := "https://www.ulipstaging.dpiit.gov.in/ulip/v1.0.0/SARATHI/01"
	reqBody := fmt.Sprintf(`{"dlnumber":"%s", "dob": "%s"}`, dlNumber, dob)

	// Make the HTTP request to the external API
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return false, err
	}

	token, err := pkg.GetAuthToken()
	if err != nil {
		return false, fmt.Errorf("failed to get auth token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var sarathiResponse models.SarathiResponse
	if err := json.NewDecoder(resp.Body).Decode(&sarathiResponse); err != nil {
		return false, err
	}

	sarathiFullName := strings.ReplaceAll(strings.ToUpper(sarathiResponse.Response[0].Response.Dldetobj[0].BioObj.BioFullName), " ", "")

	if sarathiFullName == strings.ReplaceAll(strings.ToUpper(name), " ", "") {
		return true, nil
	} else {
		return false, errors.New("DL number " + dlNumber + " is not valid")
	}
}
