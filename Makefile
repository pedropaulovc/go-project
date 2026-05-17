.PHONY: dev build run lint vet test test-coverage test-e2e test-all clean

dev:
	go run .

build:
	go build -o bin/server .

run: build
	./bin/server

lint:
	golangci-lint run ./...

vet:
	go vet ./...

test:
	go test -race ./...

test-coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...
	@echo "--- Coverage report ---"
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out -o coverage.html

test-e2e:
	npx playwright test

test-all: lint vet test-coverage test-e2e

clean:
	rm -rf bin/ coverage.out coverage.html tmp/ playwright-report/ test-results/
