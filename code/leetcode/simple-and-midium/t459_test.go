package simple_and_midium

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "abab",
			},
			want: true,
		}, {
			args: args{
				s: "abcabcabcabc",
			},
			want: true,
		}, {
			args: args{
				s: "aba",
			},
			want: false,
		}, {
			args: args{
				s: "aaaa",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repeatedSubstringPattern2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "abab",
			},
			want: true,
		}, {
			args: args{
				s: "abcabcabcabc",
			},
			want: true,
		}, {
			args: args{
				s: "aba",
			},
			want: false,
		}, {
			args: args{
				s: "aaaa",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern2(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern2() = %v, want %v", got, tt.want)
			}
		})
	}
}