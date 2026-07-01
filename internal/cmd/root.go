// Package cmd implements the CLI commands.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewRoot creates the root cobra command.
func NewRoot(version string) *cobra.Command {
	root := &cobra.Command{
		Use:     "myapp",
		Short:   "A CLI tool",
		Version: version,
	}

	root.AddCommand(newVersionCmd(version))

	return root
}

func newVersionCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if _, err := fmt.Fprintln(cmd.OutOrStdout(), version); err != nil {
				return fmt.Errorf("print version: %w", err)
			}

			return nil
		},
	}
}
