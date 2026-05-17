// Package server provides the HTTP handler and routes.
package server

import "net/http"

// New returns the application's HTTP handler with all routes registered.
func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleHome)
	mux.HandleFunc("GET /health", handleHealth)

	return mux
}

func handleHome(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(`<!DOCTYPE html>
<html><head><title>Go Project</title></head>
<body><h1>Hello, World!</h1></body></html>`))
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
