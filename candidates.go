package greenhouseio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Grayscale-Labs/greenhouse-io-go/models"
)

const (
	candidatesURL = baseURL + "/candidates?per_page=10"
)

type CandidatesRequest struct {
	client       *Client
	queryBuilder *strings.Builder
}

// Candidates returns a candidates request builder.
func (c *Client) Candidates() *CandidatesRequest {
	return &CandidatesRequest{
		client:       c,
		queryBuilder: &strings.Builder{},
	}
}

// Fetch gets a slice of candidates using the built query params.
func (r *CandidatesRequest) Fetch() ([]*models.Candidate, error) {
	//nolint:bodyclose // The fetchCandidates function will close the body.
	candidates, _, err := r.client.fetchCandidates(candidatesURL + r.queryBuilder.String())

	return candidates, err
}

func (r *CandidatesRequest) Stream(consumer chan models.Candidate, closeSignal chan error) {
	currentURL := candidatesURL + r.queryBuilder.String()

	for currentURL != "" {
		candidates, res, err := r.client.fetchCandidates(currentURL)
		if err != nil {
			closeSignal <- err

			close(consumer)
			close(closeSignal)

			break
		}

		for _, candidate := range candidates {
			consumer <- *candidate
		}

		nextURL, err := parseNextPageLink(res)
		if err != nil {
			closeSignal <- nil

			close(consumer)
			close(closeSignal)

			break
		}

		currentURL = nextURL
	}
}

func (r *CandidatesRequest) CreatedBefore(timestamp time.Time) *CandidatesRequest {
	r.addPrefixToken()

	r.queryBuilder.WriteString(fmt.Sprintf(
		"created_before=%s",
		timestamp.Format(time.RFC3339),
	))

	return r
}

// addPrefixToken adds the correct preceding token given how many parameters exist.
func (r *CandidatesRequest) addPrefixToken() {
	if r.queryBuilder.Len() == 0 {
		r.queryBuilder.WriteString("?")
	} else {
		r.queryBuilder.WriteString("&")
	}
}

// fetchCandidates fetches candidates from the given URL.
func (c *Client) fetchCandidates(url string) ([]*models.Candidate, *http.Response, error) {
	// Create request with built params string.
	req, err := c.generateHTTPRequest("GET", url)
	if err != nil {
		return nil, nil, fmt.Errorf("generating request: %w", err)
	}

	// Make request.
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("making request: %w", err)
	}

	// Defer body close.
	defer res.Body.Close()

	// Error on non-OK status.
	if res.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("status code %v", res.StatusCode)
	}

	// Read body into slice of bytes.
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("reading body: %w", err)
	}

	// Parse bytes as JSON into slice of candidates.
	var candidates []*models.Candidate
	if err := json.Unmarshal(data, &candidates); err != nil {
		return nil, nil, fmt.Errorf("unmarshaling response: %w", err)
	}

	return candidates, res, nil
}
