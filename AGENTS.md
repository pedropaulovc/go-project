# Go Project — Agent Reference

## Key Commands

```bash
make test-all   # REQUIRED before every PR (lint + vet + coverage)
make fmt        # Format (gofmt + goimports)
make tools      # Install dev tools (air, golangci-lint, goimports)
make build      # Build binary
make test       # Unit tests with race detector
```

## Structure

```
cmd/myapp/       # CLI entrypoint (main.go + run())
internal/cmd/    # Cobra command definitions
```

## Conventions

- Go 1.26+; cobra for CLI commands
- `snake_case.go` files; package names lowercase, no underscores
- No `_` for error returns; wrap errors: `fmt.Errorf("op: %w", err)`
- Structured logging via `log/slog`; no global mutable state
- Table-driven tests; coverage threshold 70%; `go test -race ./...` must pass

## Git

- Merge only — squash/rebase disabled
- Rebase to update: `git pull --rebase origin main`
- Never `--no-verify`
