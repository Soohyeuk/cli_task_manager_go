.PHONY: build test clean run fmt vet lint test-coverage help

# Binary name
BINARY_NAME=task

help:
	@echo "Available commands:"
	@echo "  make build         - Build the application"
	@echo "  make test         - Run all tests"
	@echo "  make clean        - Clean build files"
	@echo "  make run          - Run the application"
	@echo "  make fmt          - Format all Go files"
	@echo "  make vet          - Run Go vet"
	@echo "  make lint         - Run linter"
	@echo "  make test-coverage - Run tests with coverage"

build:
	go build -o $(BINARY_NAME) ./cmd/cli

test:
	go test -v ./...

clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f coverage.out

run:
	go run ./cmd/cli

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out 