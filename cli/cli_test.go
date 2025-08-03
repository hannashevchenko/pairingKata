package cli

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantInput  string
		wantErrStr string
	}{
		{name: "number input", input: "123", wantInput: "123", wantErrStr: ""},
		{name: "string input", input: "hello", wantInput: "hello", wantErrStr: ""},
		{name: "user exited", input: "", wantInput: "", wantErrStr: "user exited"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originStdin := os.Stdin
			defer func() { os.Stdin = originStdin }()
			os.Stdin = &os.File{}
			r, w, err := os.Pipe()
			if err != nil {
				log.Fatal(err)
			}
			os.Stdin = r
			go func() {
				_, err := w.WriteString(tt.input)
				if err != nil {
					t.Errorf("Failed to write to pipe: %v", err)
				}
				err = w.Close()
				if err != nil {
					return
				}
			}()
			gotInput, gotErr := Cli()
			if gotInput != tt.wantInput {
				t.Errorf("Cli() gotInput = %v, want %v", gotInput, tt.wantInput)
			}
			// Check the error
			if tt.wantErrStr == "" {
				if gotErr != nil {
					t.Errorf("Cli() returned error %v, expected nil", gotErr)
				}
			} else {
				if gotErr == nil || !strings.Contains(gotErr.Error(), tt.wantErrStr) {
					t.Errorf("Cli() returned error %v, expected error containing %q", gotErr, tt.wantErrStr)
				}
			}
		})
	}

}
