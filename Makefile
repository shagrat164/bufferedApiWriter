BINARY_NAME := "./bin/bufApiWriter"

all: lint test build

build:
	go build -v -o $(BINARY_NAME) ./cmd/bufApiWriter

run: build
	$(BINARY_NAME)

test:
	go test -race -count=10 ./...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.64.5

lint: install-lint-deps
	golangci-lint run ./...

clean:
	rm -rf ./bin

.PHONY: all build run test lint clean