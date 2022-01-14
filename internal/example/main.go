package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	greenhouseio "github.com/Grayscale-Labs/greenhouse-io-go"
)

const (
	httpTimeout = 10 * time.Second
)

func main() {
	client, err := greenhouseio.NewClient(
		context.Background(),
		&http.Client{
			Timeout: httpTimeout,
		},
		os.Getenv("GREENHOUSE_API_KEY"),
	)
	if err != nil {
		log.Fatalf("creating greenhouse client: %v", err)
	}

	candidates, err := client.Candidates().Fetch()
	if err != nil {
		log.Fatalf("fetching candidates: %v", err)
	}

	for _, c := range candidates {
		log.Default().Println(c.ID)
	}
}
