.DEFAULT_GOAL := build

.PHONY: lint
lint:
	golangci-lint run

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

.PHONY: copilot-test
copilot-test: export GH_TOKEN=$(shell gh auth token)
copilot-test: lint test
