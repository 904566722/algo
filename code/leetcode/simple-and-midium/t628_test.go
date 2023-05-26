package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{-1, 2, 3},
			},
			want: -6,
		},
		{
			args: args{
				nums: []int{-1, -2, 3},
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 24,
		},
		{
			args: args{
				nums: []int{-5, -4, -3, -2, 0, 1, 2},
			},
			want: 40,
		},
		{
			args: args{
				nums: []int{-5, -4, -3, -2, 0},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{-5, -4, -3, -2, 0, 5, 6},
			},
			want: 120,
		},
		{
			args: args{
				nums: []int{-5, -4, -3, -2, 1, 5, 6},
			},
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumProduct(tt.args.nums), "maximumProduct(%v)", tt.args.nums)
		})
	}
}
