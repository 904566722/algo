package simple_and_midium

import (
	"testing"
)


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