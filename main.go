// Package main is the entrypoint for the Go project server.
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pedropaulovc/go-project/internal/server"
)

const readHeaderTimeout = 5 * time.Second

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

	log.Printf("Server starting on :%s", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
