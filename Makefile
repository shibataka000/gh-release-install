.PHONY: fmt lint test build install clean vulncheck
.DEFAULT_GOAL := build

fmt:
	go fmt ./...
	go tool goimports -w $(shell find . -type f -name "*.go")

lint:
	go tool golangci-lint run

test:
	go test ./...

build:
	go build

install:
	go install

clean:
	go clean -testcache

vulncheck:
	go tool govulncheck ./...
