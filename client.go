package greenhouseio

import (
	"context"
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
