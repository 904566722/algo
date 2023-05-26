package simple_and_midium

import (
	"testing"
)

func Test_longestStrChain(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				words: []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestStrChain(tt.args.words); got != tt.want {
				t.Errorf("longestStrChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isQianshen(t *testing.T) {
	type args struct {
		short string
		long  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				short: "ab",
				long: "abc",
			},
			want: true,
		},{
			args: args{
				short: "ab",
				long: "acb",
			},
			want: true,
		},{
			args: args{
				short: "ab",
				long: "cab",
			},
			want: true,
		},{
			args: args{
				short: "ab",
				long: "abcd",
			},
			want: false,
		},{
			args: args{
				short: "abc",
				long: "abdd",
			},
			want: false,
		},{
			args: args{
				short: "abc",
				long: "bbbc",
			},
			want: false,
		},{
			args: args{
				short: "abc",
				long: "aefc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isQianshen(tt.args.short, tt.args.long); got != tt.want {
				t.Errorf("isQianshen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestStrChain1(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				words: []string{"xbc","pcxbcf","xb","cxbc","pcxbc"},
			},
			want: 5,
		},
		{
			args: args{
				words: []string{"a","b","ba","bca","bda","bdca"},
			},
			want: 4,
		},{
			args: args{
				words: []string{"abcd","dbqca"},
			},
			want: 1,
		},{
			args: args{
				words: []string{"abc","abcd"},
			},
			want: 2,
		},{
			args: args{
				words: []string{"abc","abcd", "abce", "abcef", "dabcef", "ddbcefa"},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestStrChain(tt.args.words); got != tt.want {
				t.Errorf("longestStrChain() = %v, want %v", got, tt.want)
			}
		})
	}
}