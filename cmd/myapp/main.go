// Package main is the entrypoint for the CLI.
package main

import (
	"os"

	"github.com/pedropaulovc/go-project/internal/cmd"
)

var version = "dev" //nolint:gochecknoglobals,nolintlint // injected at build time via -ldflags

func main() {
	root := cmd.NewRoot(version)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
