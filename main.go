// Package main is the entrypoint for the Go project server.
package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/pedropaulovc/go-project/internal/server"
)

const readHeaderTimeout = 5 * time.Second

var version = "dev" //nolint:gochecknoglobals,nolintlint // injected at build time via -ldflags

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           server.New(),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	slog.Info("server starting", "port", port, "version", version)

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
