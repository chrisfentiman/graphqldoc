package parser

import "testing"

func Test_firstToUpper(t *testing.T) {
	type args struct {
		str string
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
			if got := firstToUpper(tt.args.str); got != tt.want {
				t.Errorf("firstToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstToLower(t *testing.T) {
	type args struct {
		str string
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
			if got := firstToLower(tt.args.str); got != tt.want {
				t.Errorf("firstToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_title(t *testing.T) {
	type args struct {
		input string
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
			if got := title(tt.args.input); got != tt.want {
				t.Errorf("title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runeMap(t *testing.T) {
	type args struct {
		str         string
		replaceWith []rune
		force       bool
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
			if got := runeMap(tt.args.str, tt.args.replaceWith, tt.args.force); got != tt.want {
				t.Errorf("runeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_emptyRune(t *testing.T) {
	type args struct {
		r []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := emptyRune(tt.args.r); got != tt.want {
				t.Errorf("emptyRune() = %v, want %v", got, tt.want)
			}
		})
	}
}
