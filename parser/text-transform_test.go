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
		{
			name: "Should Capitialize First letter Only",
			args: args{
				str: "welcome to the party",
			},
			want: "Welcome to the party",
		},
		{
			name: "Capitalized letter should stay capitalized",
			args: args{
				str: "Welcome to the party",
			},
			want: "Welcome to the party",
		},
		{
			name: "Shouldn't change the case of any other letters",
			args: args{
				str: "weLcOmE tO tHe PaRtY",
			},
			want: "WeLcOmE tO tHe PaRtY",
		},
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
		{
			name: "Should Lower First letter Only",
			args: args{
				str: "Welcome to the party",
			},
			want: "welcome to the party",
		},
		{
			name: "Lowercased letter should stay lowercase",
			args: args{
				str: "welcome to the party",
			},
			want: "welcome to the party",
		},
		{
			name: "Shouldn't change the case of any other letters",
			args: args{
				str: "WeLcOmE tO tHe PaRtY",
			},
			want: "weLcOmE tO tHe PaRtY",
		},
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
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should Capitalize Correct Letters",
			args: args{
				str: "welcome to the party",
			},
			want: "Welcome to the Party",
		},
		{
			name: "Capitalized Letters should stay the same",
			args: args{
				str: "Welcome to the Party",
			},
			want: "Welcome to the Party",
		},
		{
			name: "Should lower case non title words",
			args: args{
				str: "Welcome To The Party",
			},
			want: "Welcome to the Party",
		},
		{
			name: "Should change on all words unless to match title",
			args: args{
				str: "WELCOME TO THE PARTY",
			},
			want: "Welcome to the Party",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := title(tt.args.str); got != tt.want {
				t.Errorf("title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runeMap(t *testing.T) {
	type args struct {
		str         string
		replaceWith []rune
		trim        bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should replace all hyphens with spaces",
			args: args{
				str:         "welcome-to-the-party",
				replaceWith: []rune{},
				trim:        false,
			},
			want: "welcome to the party",
		},
		{
			name: "Should replace all underscores with spaces",
			args: args{
				str:         "welcome_to_the_party",
				replaceWith: []rune{},
				trim:        false,
			},
			want: "welcome to the party",
		},
		{
			name: "Should replace all underscores and remove spaces",
			args: args{
				str:         "_welcome_to_the_party_",
				replaceWith: []rune{},
				trim:        true,
			},
			want: "welcometotheparty",
		},
		{
			name: "Should replace all hyphens and remove spaces",
			args: args{
				str:         "--welcome-to-the--party-",
				replaceWith: []rune{},
				trim:        true,
			},
			want: "welcometotheparty",
		},
		{
			name: "Should replace all hyphens, normalize spacing, and keep case",
			args: args{
				str:         "-WeLcoMe-To-ThE--pArTy-",
				replaceWith: []rune{},
				trim:        false,
			},
			want: " WeLcoMe To ThE pArTy ",
		},
		{
			name: "Should replace all underscores, normalize spacing, and keep case",
			args: args{
				str:         "_WeLcoMe_To_ThE__pArTy__",
				replaceWith: []rune{},
				trim:        false,
			},
			want: " WeLcoMe To ThE pArTy ",
		},
		{
			name: "Should replace all spaces and keep case",
			args: args{
				str:         "WeLcoMe To ThE pArTy",
				replaceWith: []rune{},
				trim:        true,
			},
			want: "WeLcoMeToThEpArTy",
		},
		{
			name: "Should replace all spaces with hyphens, keep case, and remove any whitespace",
			args: args{
				str:         "WeLcoMe To ThE pArTy 32",
				replaceWith: []rune("-"),
				trim:        false,
			},
			want: "WeLcoMe-To-ThE-pArTy-32",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runeMap(tt.args.str, tt.args.replaceWith, tt.args.trim); got != tt.want {
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
		{
			name: "Should return false when rune is not empty",
			args: args{
				r: []rune("-"),
			},
			want: false,
		},
		{
			name: "Should return true when rune is empty",
			args: args{
				r: []rune{},
			},
			want: true,
		},
		{
			name: "Should return true when rune is empty",
			args: args{
				r: []rune(""),
			},
			want: true,
		},
		{
			name: "Should return false when rune is escaped",
			args: args{
				r: []rune("\\n"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := emptyRune(tt.args.r); got != tt.want {
				t.Errorf("emptyRune() = %v, want %v", got, tt.want)
			}
		})
	}
}
