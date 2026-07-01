package cmd_test

import (
	"bytes"
	"testing"

	"github.com/pedropaulovc/go-project/internal/cmd"
)

// writeError is a test-only error type that simulates a write failure.
type writeError struct{}

func (writeError) Error() string { return "broken pipe" }

// failWriter is an io.Writer that always returns an error.
type failWriter struct{}

func (failWriter) Write(_ []byte) (int, error) { return 0, writeError{} }

func TestRootExecutes(t *testing.T) {
	t.Parallel()

	root := cmd.NewRoot("test-version")
	root.SetArgs([]string{"--help"})

	buf := &bytes.Buffer{}
	root.SetOut(buf)

	if err := root.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("expected help output, got empty")
	}
}

func TestVersionCommand(t *testing.T) {
	t.Parallel()

	root := cmd.NewRoot("v1.2.3")
	root.SetArgs([]string{"version"})

	buf := &bytes.Buffer{}
	root.SetOut(buf)

	if err := root.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := buf.String(); got != "v1.2.3\n" {
		t.Errorf("expected 'v1.2.3\\n', got %q", got)
	}
}

func TestVersionCommandPropagatesWriteError(t *testing.T) {
	t.Parallel()

	root := cmd.NewRoot("v1.2.3")
	root.SetArgs([]string{"version"})
	root.SetOut(&failWriter{})
	root.SilenceErrors = true
	root.SilenceUsage = true

	if err := root.Execute(); err == nil {
		t.Fatal("expected error from write failure, got nil")
	}
}
