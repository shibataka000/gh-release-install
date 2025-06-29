.DEFAULT_GOAL := build

.PHONY: lint
lint:
	go tool golangci-lint run

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build

.PHONY: install
install:
	go install

.PHONY: clean
clean:
	go clean -testcache
