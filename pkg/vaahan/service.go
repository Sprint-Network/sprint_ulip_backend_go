package vaahan_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"prechecks/pkg"
)

func FetchVehicleDetailsFromULIP(vehicleNumber string) (map[string]interface{}, error) {
	url := "https://www.ulipstaging.dpiit.gov.in/ulip/v1.0.0/VAHAN/01"
	reqBody := fmt.Sprintf(`{"vehiclenumber":"%s"}`, vehicleNumber)

	// Make the HTTP request to the external API
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return nil, err
	}

	token, err := pkg.GetAuthToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get auth token: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response body
	var apiResponse struct {
		Response []struct {
			Response string `json:"response"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	if len(apiResponse.Response) == 0 {
		return nil, fmt.Errorf("no vehicle details found")
	}

	// Extract the XML response from the API and convert to JSON
	return ConvertXMLToJSON(apiResponse.Response[0].Response)
}
