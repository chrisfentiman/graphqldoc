package parser

import "testing"

func TestHTTP(t *testing.T) {
	type args struct {
		endpoint  string
		headers   []string
		templates string
		format    bool
		overwrite bool
		out       string
		dryRun    bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HTTP(tt.args.endpoint, tt.args.headers, tt.args.templates, tt.args.format, tt.args.overwrite, tt.args.out, tt.args.dryRun)
		})
	}
}
