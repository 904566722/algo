package simple_and_midium

import "testing"

func Test_isVowel(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				b: 'a',
			},
			want: true,
		}, {
			args: args{
				b: 'b',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVowel(tt.args.b); got != tt.want {
				t.Errorf("isVowel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "hello",
			},
			want: "holle",
		},
		{
			args: args{
				s: "leetcode",
			},
			want: "leotcede",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
