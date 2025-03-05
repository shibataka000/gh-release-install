.PHONY: lint test build install clean
.DEFAULT_GOAL := build

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
