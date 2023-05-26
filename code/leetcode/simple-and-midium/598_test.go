package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxCount(t *testing.T) {
	type args struct {
		m   int
		n   int
		ops [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				m: 3,
				n: 3,
				ops: [][]int{
					{2, 2},
					{3, 3},
				},
			},
			want: 4,
		},
		{
			args: args{
				m:   3,
				n:   3,
				ops: [][]int{},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxCount(tt.args.m, tt.args.n, tt.args.ops), "maxCount(%v, %v, %v)", tt.args.m, tt.args.n, tt.args.ops)
		})
	}
}
