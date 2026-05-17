# Agent Standards

## Exit Criteria

All PRs must pass: `make test-all` (lint + unit tests with coverage + E2E tests).

## Code Conventions

### Go

- Go 1.26+ required
- Use `net/http` standard library (no frameworks unless justified)
- Errors must be handled explicitly — never use `_` for error returns
- Use structured logging (`log/slog`)
- Prefer table-driven tests
- No global mutable state

### File Naming

- `snake_case.go` for all Go files
- `snake_case_test.go` for test files
- Package names: short, lowercase, no underscores

### Error Handling

- Wrap errors with context: `fmt.Errorf("operation: %w", err)`
- Let errors propagate to appropriate boundaries
- Validate at system boundaries (HTTP handlers, CLI input)

### Testing

- `go test ./...` must pass with zero failures
- Coverage threshold: 70%
- New features require: unit tests + E2E tests
- No flaky tests — fix immediately
- E2E: max 5s per test, 1min suite timeout
- Prefer table-driven tests with subtests

### Git Workflow

- Merge only (`gh pr merge --merge --auto`)
- Squash/rebase merge disabled
- PRs must be up-to-date with main before merging
- Rebase to update: `git pull --rebase origin main`
- Never bypass hooks (`--no-verify`)

## Multi-Instance Port Management

Worktree-based port mapping:
- Worktree A=8010, B=8020, C=8030, D=8040, E=8050, F=8060, G=8070
- Non-worktree: 8080 (default)

Set via `PORT` environment variable.
