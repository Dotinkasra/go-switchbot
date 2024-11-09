package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type apiResponse struct {
	StatusCode int    `json:"statusCode"`
	Body       any    `json:"body"`
	Message    string `json:"message"`
}

// GetToken: Reads the environment variable “SWITCHBOT_TOKEN” and returns a token.
func GetToken() (string, error) {
	token := os.Getenv("SWITCHBOT_TOKEN")
	if token == "" {
		return "", fmt.Errorf("don't exists setting file")
	}
	return token, nil
}

type switchbotAPIClient struct {
	AuthToken string
}

// NewSwitchbotAPIClient: Create a client to send to Switchbot's API.
func NewSwitchbotAPIClient() *switchbotAPIClient {
	token, err := GetToken()
	if err != nil {
		panic(err)
	}
	return &switchbotAPIClient{AuthToken: token}
}

// SendAPIRequest: Send to Switchbot's API.
func (c *switchbotAPIClient) SendAPIRequest(uri string, method string, payload interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf8")
	req.Header.Set("Authorization", c.AuthToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == 200) {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}
	return body, nil
}
