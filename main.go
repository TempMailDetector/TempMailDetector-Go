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
	apiURL      = "https://api.tempmaildetector.com/check"
	contentType = "application/json"
)

type Client struct {
	APIKey string
}

type DomainCheckRequest struct {
	Domain string `json:"domain"`
}

type DomainCheckResponse struct {
	Domain string `json:"domain"`
	Score  int    `json:"score"`
	Meta   struct {
		BlockList           bool `json:"block_list"`
		DomainAge           int  `json:"domain_age"`
		WebsiteResolves     bool `json:"website_resolves"`
		AcceptsAllAddresses bool `json:"accepts_all_addresses"`
		ValidEmailSecurity  bool `json:"valid_email_security"`
	} `json:"meta"`
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) CheckDomain(domain string) (*DomainCheckResponse, error) {
	requestBody, err := json.Marshal(DomainCheckRequest{Domain: domain})
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

	var response DomainCheckResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
