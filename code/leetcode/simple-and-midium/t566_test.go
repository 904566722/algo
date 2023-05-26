package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	type args struct {
		mat [][]int
		r   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				mat: [][]int{
					{1, 2},
					{3, 4},
				},
				r: 1,
				c: 4,
			},
			want: [][]int{
				{1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, matrixReshape(tt.args.mat, tt.args.r, tt.args.c), "matrixReshape(%v, %v, %v)", tt.args.mat, tt.args.r, tt.args.c)
		})
	}
}

func Test_matrixReshape1(t *testing.T) {
	type args struct {
		mat [][]int
		r   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				mat: [][]int{
					{1, 2},
					{3, 4},
				},
				r: 1,
				c: 4,
			},
			want: [][]int{
				{1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, matrixReshape(tt.args.mat, tt.args.r, tt.args.c), "matrixReshape(%v, %v, %v)", tt.args.mat, tt.args.r, tt.args.c)
		})
	}
}

func Test_matrixReshape2(t *testing.T) {
	type args struct {
		mat [][]int
		r   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				mat: [][]int{
					{1, 2},
					{3, 4},
				},
				r: 1,
				c: 4,
			},
			want: [][]int{
				{1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, matrixReshape2(tt.args.mat, tt.args.r, tt.args.c), "matrixReshape2(%v, %v, %v)", tt.args.mat, tt.args.r, tt.args.c)
		})
	}
}
