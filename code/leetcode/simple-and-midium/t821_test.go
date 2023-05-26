package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shortestToChar(t *testing.T) {
	type args struct {
		s string
		c byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				s: "loveleetcode",
				c: 'e',
			},
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, shortestToChar(tt.args.s, tt.args.c), "shortestToChar(%v, %v)", tt.args.s, tt.args.c)
		})
	}
}
