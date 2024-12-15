package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	Client http.Client
}

func (c *Client) ServerPostRequest(port string, testData interface{}, endpoint string) (*http.Response, error) {
	payloadBytes, err := json.Marshal(testData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%s/%s", port, endpoint), body)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	c.Client.Timeout = 30 * time.Second
	return c.getResponse(req)
}

func (c *Client) ServerGetRequest(port string, userID string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%s/user/%s", port, userID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	c.Client.Timeout = 30 * time.Second

	return c.getResponse(req)
}

func (c *Client) getResponse(req *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("response err: %w", err)
	}

	return resp, nil
}
