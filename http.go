package greenhouseio

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

const (
	baseURL = "https://harvest.greenhouse.io/v1"
)

var (
	nextPageLinkRegex = regexp.MustCompile(`<(.*)>; rel="next"`)
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
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

// parseNextPageLink parses the `link` header in the given response for the `next` link.
func parseNextPageLink(res *http.Response) (string, error) {
	matches := nextPageLinkRegex.FindStringSubmatch((res.Header.Get("link")))
	if len(matches) == 0 {
		return "", errors.New("no next link")
	}

	// First element is the match, second element is the group.
	return matches[1], nil
}
