package parser

import (
	"reflect"
	"testing"
)

func Test_outFiles(t *testing.T) {
	type args struct {
		out string
	}
	tests := []struct {
		name string
		args args
		want *gqlFiles
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outFiles(tt.args.out); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_docGenerator_mkdir(t *testing.T) {
	tests := []struct {
		name    string
		d       *docGenerator
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.mkdir(); (err != nil) != tt.wantErr {
				t.Errorf("docGenerator.mkdir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_relativePath(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relativePath(tt.args.dir); got != tt.want {
				t.Errorf("relativePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absolutePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := absolutePath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("absolutePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("absolutePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
