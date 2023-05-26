package simple_and_midium

import (
	"reflect"
	"testing"
)

func Test_getNext(t *testing.T) {
	type args struct {
		pat string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				pat: "aabaacaad",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNext(tt.args.pat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kmp(t *testing.T) {
	type args struct {
		s   string
		pat string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "abcdefg",
				pat: "cde",
			},
			want: true,
		},{
			args: args{
				s: "abcdefg",
				pat: "cdf",
			},
			want: false,
		},{
			args: args{
				s: "abcdefg",
				pat: "abcdefg",
			},
			want: true,
		},{
			args: args{
				s: "abcdefg",
				pat: "abc",
			},
			want: true,
		},{
			args: args{
				s: "abcdefg",
				pat: "efg",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmp(tt.args.s, tt.args.pat); got != tt.want {
				t.Errorf("kmp() = %v, want %v", got, tt.want)
			}
		})
	}
}