package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	authToken       string
	authTokenMutex  sync.Mutex
	authTokenExpiry time.Time
)

// FetchAuthToken retrieves a new token and updates the cache.
func FetchAuthToken() (string, error) {
	apiURL := "https://www.ulipstaging.dpiit.gov.in/ulip/v1.0.0/user/login"
	payload := map[string]string{
		"username": "<username>",
		"password": "<password>",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to serialize login payload: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create login request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("login request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read login response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login API returned status %d: %s", resp.StatusCode, string(body))
	}

	var apiResponse struct {
		Response struct {
			ID string `json:"id"`
		} `json:"response"`
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return "", fmt.Errorf("failed to parse login response JSON: %w", err)
	}

	if apiResponse.Response.ID == "" {
		return "", fmt.Errorf("auth token not found in login response")
	}

	// Update global token and expiry
	authTokenMutex.Lock()
	defer authTokenMutex.Unlock()
	authToken = apiResponse.Response.ID

	// Set expiry to 29 minutes from now (since the token is valid for 30 minutes)
	authTokenExpiry = time.Now().Add(29 * time.Minute)

	return authToken, nil
}

// GetAuthToken safely retrieves a valid token, refreshing if expired.
func GetAuthToken() (string, error) {
	//authTokenMutex.Lock()
	//defer authTokenMutex.Unlock()

	// Initialize authTokenExpiry the first time it's needed
	if authTokenExpiry.IsZero() {
		// Token is being used for the first time, fetch it
		return FetchAuthToken()
	}

	// Check if the token has expired, and fetch a new one if it has
	if time.Now().After(authTokenExpiry) {
		return FetchAuthToken()
	}

	// Return the existing token if it's still valid
	return authToken, nil
}
