# Go Project

Template project for Go HTTP applications.

## Requirements

- Go >= 1.24
- Node >= 24 (for E2E tests)
- golangci-lint

## Quick Start

```bash
make dev          # Run dev server on :8080
make build        # Build binary to bin/server
make test-all     # Run lint + tests + E2E
```

## Scripts

| Command            | Description                          |
| ------------------ | ------------------------------------ |
| `make dev`         | Run development server               |
| `make build`       | Build production binary               |
| `make run`         | Build and run                         |
| `make lint`        | Run golangci-lint                     |
| `make test`        | Run unit tests                        |
| `make test-coverage` | Run tests with coverage report     |
| `make test-e2e`    | Run Playwright E2E tests              |
| `make test-all`    | Full suite: lint + coverage + E2E     |
| `make clean`       | Remove build artifacts                |

## Project Structure

```
.
├── main.go                 # Entrypoint
├── internal/
│   └── server/             # HTTP server & handlers
├── e2e/                    # Playwright E2E tests
├── .github/workflows/      # CI/CD pipelines
├── .husky/                 # Git hooks
└── Makefile                # Build & dev commands
```
