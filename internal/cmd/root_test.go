package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/pedropaulovc/go-project/internal/cmd"
)

func TestRootCommand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		args        []string
		version     string
		wantErr     bool
		wantContain string
	}{
		{
			name:        "help flag prints usage",
			args:        []string{"--help"},
			version:     "test-version",
			wantContain: "myapp",
		},
		{
			name:        "version flag prints version",
			args:        []string{"--version"},
			version:     "v1.2.3",
			wantContain: "v1.2.3",
		},
		{
			name:        "version subcommand prints version",
			args:        []string{"version"},
			version:     "v1.2.3",
			wantContain: "v1.2.3\n",
		},
		{
			name:    "unknown command returns error",
			args:    []string{"notacommand"},
			version: "test-version",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			root := cmd.NewRoot(tc.version)
			root.SetArgs(tc.args)

			buf := &bytes.Buffer{}
			root.SetOut(buf)
			root.SetErr(buf)

			err := root.Execute()
			if (err != nil) != tc.wantErr {
				t.Fatalf("Execute() error = %v, wantErr %v", err, tc.wantErr)
			}

			if tc.wantContain != "" && !strings.Contains(buf.String(), tc.wantContain) {
				t.Errorf("output %q does not contain %q", buf.String(), tc.wantContain)
			}
		})
	}
}
