.PHONY: dev build run lint test test-coverage test-e2e test-all clean

dev:
	go run .

build:
	go build -o bin/server .

run: build
	./bin/server

lint:
	golangci-lint run ./...

test:
	go test ./...

test-coverage:
	go test -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

test-e2e:
	npx playwright test

test-all: lint test-coverage test-e2e

clean:
	rm -rf bin/ coverage.out coverage.html tmp/ playwright-report/ test-results/
