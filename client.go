// Package upbit provides a Go SDK for the Upbit cryptocurrency exchange API.
package upbit

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	BaseURL = "https://api.upbit.com/v1"
)

// Client is the main Upbit API client.
type Client struct {
	accessKey  string
	secretKey  string
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new Upbit API client.
// For public API endpoints, you can pass empty strings for accessKey and secretKey.
func NewClient(accessKey, secretKey string) *Client {
	return &Client{
		accessKey: accessKey,
		secretKey: secretKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: BaseURL,
	}
}

// SetHTTPClient allows you to set a custom HTTP client.
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

// SetBaseURL allows you to set a custom base URL (useful for testing).
func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

// generateToken creates a JWT token for authenticated requests.
func (c *Client) generateToken(queryParams url.Values) (string, error) {
	claims := jwt.MapClaims{
		"access_key": c.accessKey,
		"nonce":      uuid.New().String(),
		"timestamp":  time.Now().UnixMilli(),
	}

	if len(queryParams) > 0 {
		queryString := queryParams.Encode()
		hash := sha512.Sum512([]byte(queryString))
		claims["query_hash"] = hex.EncodeToString(hash[:])
		claims["query_hash_alg"] = "SHA512"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.secretKey))
}

// doRequest performs an HTTP request.
func (c *Client) doRequest(method, endpoint string, params url.Values, authenticated bool) ([]byte, error) {
	urlStr := c.baseURL + endpoint
	var body io.Reader

	if method == http.MethodGet && len(params) > 0 {
		urlStr += "?" + params.Encode()
	} else if len(params) > 0 {
		body = strings.NewReader(params.Encode())
	}

	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	if method == http.MethodPost || method == http.MethodDelete {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if authenticated {
		token, err := c.generateToken(params)
		if err != nil {
			return nil, fmt.Errorf("failed to generate token: %w", err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := json.Unmarshal(respBody, &apiErr); err != nil {
			return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
		}
		return nil, &apiErr
	}

	return respBody, nil
}

// get performs a GET request.
func (c *Client) get(endpoint string, params url.Values, authenticated bool) ([]byte, error) {
	return c.doRequest(http.MethodGet, endpoint, params, authenticated)
}

// post performs a POST request.
func (c *Client) post(endpoint string, params url.Values, authenticated bool) ([]byte, error) {
	return c.doRequest(http.MethodPost, endpoint, params, authenticated)
}

// delete performs a DELETE request.
func (c *Client) delete(endpoint string, params url.Values, authenticated bool) ([]byte, error) {
	return c.doRequest(http.MethodDelete, endpoint, params, authenticated)
}
