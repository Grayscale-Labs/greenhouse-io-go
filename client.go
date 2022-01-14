package greenhouseio

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://harvest.greenhouse.io/v1"
)

type Client struct {
	ctx context.Context

	httpClient      httpClient
	authHeaderValue string
}

func NewClient(ctx context.Context, httpClient httpClient, apiToken string) (*Client, error) {
	return &Client{
		ctx:             ctx,
		httpClient:      httpClient,
		authHeaderValue: generateAuthHeaderValue(apiToken),
	}, nil
}

// generateHTTPRequest creates a request object with the auth header and client context attached.
func (c *Client) generateHTTPRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(c.ctx, method, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Authorization", c.authHeaderValue)

	return req, nil
}

// generateAuthHeaderValue encodes the API token and formats it according to RFC 7617.
func generateAuthHeaderValue(apiToken string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(apiToken + ":"))

	return "Basic " + encoded
}
