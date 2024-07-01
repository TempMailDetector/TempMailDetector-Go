package tempmaildetector

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiURL      = "https://api.tempmaildetector.com/check-email"
	contentType = "application/json"
)

type Client struct {
	APIKey string
}

type EmailCheckRequest struct {
	Email string `json:"email"`
}

type EmailCheckResponse struct {
	Email string `json:"email"`
	Score int    `json:"score"`
	Meta  struct {
		BlockList           bool `json:"block_list"`
		DomainAge           int  `json:"domain_age"`
		WebsiteResolves     bool `json:"website_resolves"`
		RandomCharacters    bool `json:"random_characters"`
		AcceptsAllAddresses bool `json:"accepts_all_addresses"`
		UsesPlus            bool `json:"uses_plus"`
	} `json:"meta"`
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) CheckEmail(email string) (*EmailCheckResponse, error) {
	requestBody, err := json.Marshal(EmailCheckRequest{Email: email})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("received non-200 response: " + string(body))
	}

	var response EmailCheckResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
