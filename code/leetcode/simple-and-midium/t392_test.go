package simple_and_midium

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		}, {
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSubsequence2(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		}, {
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		}, {
			args: args{
				s: "a",
				t: "ahbgdc",
			},
			want: true,
		}, {
			args: args{
				s: "a",
				t: "bbbbbbb",
			},
			want: false,
		}, {
			args: args{
				s: "aa",
				t: "abbbbbbb",
			},
			want: false,
		}, {
			args: args{
				s: "aa",
				t: "abbbbbbba",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence2() = %v, want %v", got, tt.want)
			}
		})
	}
}
