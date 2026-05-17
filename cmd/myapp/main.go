// Package main is the entrypoint for the CLI.
package main

import (
	"os"

	"github.com/pedropaulovc/go-project/internal/cmd"
)

var version = "dev" //nolint:gochecknoglobals,nolintlint // injected at build time via -ldflags

func main() {
	if err := run(version, os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

func run(ver string, args []string) error {
	root := cmd.NewRoot(ver)
	root.SetArgs(args)
	return root.Execute()
}
