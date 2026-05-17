#!/usr/bin/env bash
set -euo pipefail

# Idempotent repo provisioning script.
# Safe to run multiple times — skips steps that are already done.

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$REPO_ROOT"

echo "==> Provisioning go-project..."

# 1. Check Go is installed.
if ! command -v go &>/dev/null; then
  echo "ERROR: Go is not installed. Install Go 1.24+ first."
  exit 1
fi
echo "  Go: $(go version)"

# 2. Check golangci-lint is installed.
if ! command -v golangci-lint &>/dev/null; then
  echo "  Installing golangci-lint..."
  go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6
fi
echo "  golangci-lint: $(golangci-lint version --short 2>&1 || golangci-lint --version)"

# 3. Install Go dependencies.
echo "  Running go mod download..."
go mod download

# 4. Install Node dependencies (for Playwright E2E).
if command -v npm &>/dev/null; then
  if [ ! -d node_modules ]; then
    echo "  Installing npm dependencies..."
    npm ci
  else
    echo "  node_modules already present, skipping npm ci."
  fi

  # Install Playwright browsers.
  if ! npx playwright install --dry-run chromium &>/dev/null 2>&1; then
    echo "  Installing Playwright browsers..."
    npx playwright install chromium --with-deps
  else
    echo "  Playwright browsers already installed."
  fi
else
  echo "  WARN: npm not found — skipping Playwright setup. Install Node 24+ for E2E tests."
fi

# 5. Set up git hooks (Husky).
if [ -d .husky ]; then
  git config core.hooksPath .husky
  echo "  Git hooks path set to .husky"
fi

# 6. Verify the setup works.
echo ""
echo "==> Verification..."
echo "  Running lint..."
golangci-lint run ./...
echo "  Running tests..."
go test -race ./...
echo ""
echo "==> Done! Run 'make dev' to start the server."
