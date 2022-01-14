package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	greenhouseio "github.com/Grayscale-Labs/greenhouse-io-go"
	"github.com/Grayscale-Labs/greenhouse-io-go/models"
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

	candidates, closeSignal := make(chan models.Candidate), make(chan error)
	go client.Candidates().Stream(candidates, closeSignal)

	for {
		select {
		case err := <-closeSignal:
			if err != nil {
				log.Fatalf("error streaming candidates: %v", err)
			}

			os.Exit(0)
		case candidate := <-candidates:
			log.Default().Println("streamed", candidate.ID)
		}
	}
}
