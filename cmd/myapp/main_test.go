// Package main_test contains integration tests for the myapp CLI binary.
package main_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// buildBinary compiles the binary into a temp directory and returns its path.
func buildBinary(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	name := "myapp"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}

	out := filepath.Join(dir, name)
	cmd := exec.Command("go", "build", "-o", out, ".")
	cmd.Dir = filepath.Join(findModuleRoot(t), "cmd", "myapp")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v\n%s", err, output)
	}

	return out
}

// findModuleRoot walks up from the test file's directory to find go.mod.
func findModuleRoot(t *testing.T) string {
	t.Helper()

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("could not find go.mod")
		}

		dir = parent
	}
}

func TestCLIIntegration(t *testing.T) {
	t.Parallel()

	bin := buildBinary(t)

	tests := []struct {
		name       string
		args       []string
		wantExit   int
		wantStdout string
		wantStderr string
	}{
		{
			name:       "no args prints help",
			args:       []string{},
			wantExit:   0,
			wantStdout: "myapp",
		},
		{
			name:       "help flag prints usage",
			args:       []string{"--help"},
			wantExit:   0,
			wantStdout: "Usage:",
		},
		{
			name:       "version subcommand prints version",
			args:       []string{"version"},
			wantExit:   0,
			wantStdout: "dev",
		},
		{
			name:     "unknown command exits non-zero",
			args:     []string{"notacommand"},
			wantExit: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			//nolint:gosec // test binary path is constructed in this test
			cmd := exec.Command(bin, tc.args...)
			out, err := cmd.CombinedOutput()

			exitCode := 0
			if err != nil {
				if exitErr, ok := err.(*exec.ExitError); ok {
					exitCode = exitErr.ExitCode()
				} else {
					t.Fatalf("unexpected error: %v", err)
				}
			}

			if exitCode != tc.wantExit {
				t.Errorf("exit code = %d, want %d\noutput: %s", exitCode, tc.wantExit, out)
			}

			if tc.wantStdout != "" && !strings.Contains(string(out), tc.wantStdout) {
				t.Errorf("output %q does not contain %q", string(out), tc.wantStdout)
			}
		})
	}
}
