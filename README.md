# Greenhouse IO

A Go interface to [Greenhouse.io's](https://app.greenhouse.io/jobboard/jsonp_instructions) API

## Useage

### Creating the Client

`NewClient` accepts:

- A context; used for any HTTP requests made using the client.
- Any struct implementing the [`httpClient`](https://github.com/Grayscale-Labs/greenhouse-io-go/blob/55cec5d058f98725c3c071945407652f10d347de/http.go#L5) interface; this allows for useage of other HTTP clients, like (go-retryablehttp)[https://github.com/hashicorp/go-retryablehttp].
- The API token.

```
client, err := greenhouseio.NewClient(
  context.Background(),
  &http.Client{
    Timeout: httpTimeout,
  },
  os.Getenv("GREENHOUSE_API_KEY"),
)
```

### Single Request

All candidates

```
candidates, err := client.Candidates().Fetch()
```

All candidates created before

```
candidates, err := client.Candidates().CreatedBefore(time.Now()).Fetch()
```

### Streaming

Streaming puts each resource loaded into a given consumer channel while following the `next` URL found in each response header. When an error occurs or no `next` URL is found, an error or nil is sent to a given close signal channel.

All candidates

```
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
```

## API Documentation

Documentation for the Harvest and Job Board web APIs can be found at [developers.greenhouse.io](https://developers.greenhouse.io).

## Development

### Prerequisites

Installing Go and [golangci-lint](https://golangci-lint.run/) via Brew

```bash
brew install go golangci-lint
echo "export PATH=~/go/bin/:$PATH" >> ~/.zshrc
```

### Commands

Linting

```bash
make lint
```

Run tests

```bash
make test
```
