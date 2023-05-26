package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_transpose(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			want: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, transpose(tt.args.matrix), "transpose(%v)", tt.args.matrix)
		})
	}
}
