package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "help flag",
			args:    []string{"--help"},
			wantErr: false,
		},
		{
			name:    "version command",
			args:    []string{"version"},
			wantErr: false,
		},
		{
			name:    "unknown command",
			args:    []string{"unknown-command-xyz"},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := run("test-version", tc.args)
			if (err != nil) != tc.wantErr {
				t.Errorf("run(%v) error = %v, wantErr %v", tc.args, err, tc.wantErr)
			}
		})
	}
}
