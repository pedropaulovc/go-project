# Go Project

Template project for Go HTTP applications.

## Requirements

- Go >= 1.26
- Node >= 26 (for E2E tests)

## Quick Start

```bash
make tools        # Install dev tools (air, golangci-lint, goimports)
make dev          # Run dev server on :8080 with hot reload
make build        # Build binary to bin/server
make test-all     # Run lint + tests + E2E
```

## Scripts

| Command              | Description                              |
| -------------------- | ---------------------------------------- |
| `make dev`           | Run dev server with hot reload (air)     |
| `make build`         | Build production binary with version tag |
| `make run`           | Build and run                            |
| `make fmt`           | Format code (gofmt + goimports)          |
| `make lint`          | Run golangci-lint                        |
| `make vet`           | Run go vet                               |
| `make test`          | Run unit tests with race detector        |
| `make test-coverage` | Run tests with coverage report           |
| `make test-e2e`      | Run Playwright E2E tests                 |
| `make test-all`      | Full suite: lint + vet + coverage + E2E  |
| `make tools`         | Install dev tools                        |
| `make clean`         | Remove build artifacts                   |

## Project Structure

```
.
├── main.go                 # Entrypoint
├── internal/
│   └── server/             # HTTP server & handlers
├── e2e/                    # Playwright E2E tests
├── scripts/
│   ├── dev.sh              # Dev server launcher (port mapping, zombie cleanup)
│   └── provision-repo.sh   # GitHub repo provisioning (rulesets, settings)
├── .github/workflows/      # CI/CD pipelines
├── .husky/                 # Git hooks
├── .air.toml               # Hot reload config
└── Makefile                # Build & dev commands
```
