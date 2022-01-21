.PHONY: lint test

lint:
	bin/docker "golangci-lint run"

test:
	bin/docker 'go test ./...'
