package server_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pedropaulovc/go-project/internal/server"
)

func TestHandleHome(t *testing.T) {
	t.Parallel()

	srv := server.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "Hello, World!") {
		t.Errorf("expected body to contain 'Hello, World!', got %q", rec.Body.String())
	}
}

func TestHandleHealth(t *testing.T) {
	t.Parallel()

	srv := server.New()
	req := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/health", http.NoBody)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	if rec.Body.String() != `{"status":"ok"}` {
		t.Errorf("unexpected body: %q", rec.Body.String())
	}
}
