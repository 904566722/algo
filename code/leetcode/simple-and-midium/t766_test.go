package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isToeplitzMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4},
					{5, 1, 2, 3},
					{9, 5, 1, 2},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isToeplitzMatrix(tt.args.matrix), "isToeplitzMatrix(%v)", tt.args.matrix)
		})
	}
}
