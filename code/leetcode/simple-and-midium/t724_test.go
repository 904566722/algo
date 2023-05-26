package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
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
				nums: []int{1, 7, 3, 6, 5, 6},
			},
			want: 3,
		}, {
			args: args{
				nums: []int{1, 2, 3},
			},
			want: -1,
		}, {
			args: args{
				nums: []int{2, 1, -1},
			},
			want: 0,
		}, {
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, pivotIndex(tt.args.nums), "pivotIndex(%v)", tt.args.nums)
		})
	}
}

func Test_pivotIndex2(t *testing.T) {
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
				nums: []int{1, 7, 3, 6, 5, 6},
			},
			want: 3,
		}, {
			args: args{
				nums: []int{1, 2, 3},
			},
			want: -1,
		}, {
			args: args{
				nums: []int{2, 1, -1},
			},
			want: 0,
		}, {
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, pivotIndex2(tt.args.nums), "pivotIndex2(%v)", tt.args.nums)
		})
	}
}
