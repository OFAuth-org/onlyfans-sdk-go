// Package ofauth provides a minimal Go client for the OFAuth API
package ofauth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const DefaultBaseURL = "https://api-next.ofauth.com"

// Config holds the client configuration
type Config struct {
	APIKey       string
	BaseURL      string
	ConnectionID string
	HTTPClient   *http.Client
}

// Client is the OFAuth API client
type Client struct {
	config Config
}

// NewClient creates a new OFAuth client
func NewClient(apiKey string) *Client {
	return &Client{
		config: Config{
			APIKey:     apiKey,
			BaseURL:    DefaultBaseURL,
			HTTPClient: http.DefaultClient,
		},
	}
}

// WithBaseURL sets a custom base URL
func (c *Client) WithBaseURL(baseURL string) *Client {
	c.config.BaseURL = strings.TrimRight(baseURL, "/")
	return c
}

// WithConnectionID sets a default connection ID
func (c *Client) WithConnectionID(connectionID string) *Client {
	c.config.ConnectionID = connectionID
	return c
}

// WithHTTPClient sets a custom HTTP client
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.config.HTTPClient = httpClient
	return c
}

// APIError represents an API error response
type APIError struct {
	Status  int
	Message string
	Code    string
	Details interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("OFAuth API error %d: %s", e.Status, e.Message)
}

// RequestOptions holds options for a request
type RequestOptions struct {
	ConnectionID string
	Query        url.Values
	Body         interface{}
}

// Request makes an API request
func (c *Client) Request(ctx context.Context, method, path string, opts *RequestOptions) ([]byte, error) {
	reqURL := c.config.BaseURL + path
	
	if opts != nil && len(opts.Query) > 0 {
		reqURL += "?" + opts.Query.Encode()
	}
	
	var bodyReader io.Reader
	if opts != nil && opts.Body != nil {
		bodyBytes, err := json.Marshal(opts.Body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}
	
	req, err := http.NewRequestWithContext(ctx, method, reqURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	
	req.Header.Set("apiKey", c.config.APIKey)
	
	connID := c.config.ConnectionID
	if opts != nil && opts.ConnectionID != "" {
		connID = opts.ConnectionID
	}
	if connID != "" {
		req.Header.Set("x-connection-id", connID)
	}
	
	if opts != nil && opts.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	
	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()
	
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}
	
	if resp.StatusCode >= 400 {
		var apiErr struct {
			Message string      `json:"message"`
			Code    string      `json:"code"`
			Details interface{} `json:"details"`
		}
		if err := json.Unmarshal(respBody, &apiErr); err == nil {
			return nil, &APIError{
				Status:  resp.StatusCode,
				Message: apiErr.Message,
				Code:    apiErr.Code,
				Details: apiErr.Details,
			}
		}
		return nil, &APIError{
			Status:  resp.StatusCode,
			Message: string(respBody),
		}
	}
	
	return respBody, nil
}
