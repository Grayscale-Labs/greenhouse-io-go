package greenhouseio

import (
	"context"
	"net/http"
	"testing"
)

func TestGeneratedHTTPRequest(t *testing.T) {
	expectedMethod := "GET"
	expectedURL := "https://grayscaleapp.com/"
	expectedAuthHeaderVal := "reuben sandwich"

	client := &Client{
		ctx:             context.Background(),
		authHeaderValue: "reuben sandwich",
	}

	req, err := client.generateHTTPRequest("GET", "https://grayscaleapp.com/")
	if err != nil {
		t.Fatal(err)
	}

	if req.Method != expectedMethod {
		t.Errorf("unexpected method: got %v, expected %v", req.Method, expectedMethod)
	}

	if reqURL := req.URL.String(); reqURL != expectedURL {
		t.Errorf("unexpected URL: got %v, expected %v", reqURL, expectedURL)
	}

	if authHeaderVal := req.Header.Get("Authorization"); authHeaderVal != expectedAuthHeaderVal {
		t.Errorf("unexpected auth header value: got %v, expected %v", authHeaderVal, expectedAuthHeaderVal)
	}
}

func TestNextPageLinkExists(t *testing.T) {
	expectedResult := "https://harvest.greenhouse.io/v1/candidates?page=2&per_page=10"

	res := &http.Response{
		Header: http.Header{},
	}
	res.Header.Set(
		"link",
		`<https://harvest.greenhouse.io/v1/candidates?page=2&per_page=10>; rel="next",`+
			`<https://harvest.greenhouse.io/v1/candidates?page=16&per_page=10>; rel="last"`,
	)

	result := parseNextPageLink(res)
	if result != expectedResult {
		t.Errorf("unexpected result: got %v, expected %v", result, expectedResult)
	}
}

func TestNextPageLinkDoesntExist(t *testing.T) {
	res := &http.Response{
		Header: http.Header{},
	}
	res.Header.Set(
		"link",
		`<https://harvest.greenhouse.io/v1/offices?page=1&per_page=10>; rel="last"`,
	)

	result := parseNextPageLink(res)
	if result != "" {
		t.Errorf("unexpected result: got %v, expected empty string", result)
	}
}
