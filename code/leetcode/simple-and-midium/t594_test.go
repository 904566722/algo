package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findLHS(t *testing.T) {
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
				nums: []int{1, 4, 1, 3, 1, -14, 1, -13},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{1, 2, 1, 3, 0, 0, 2, 2, 1, 3, 3},
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{-3, -1, -1, -1, -3, -2},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 3, 5, 7, 9, 11, 13, 15, 17},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{1, 3, 2, 2, 5, 2, 3, 7},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findLHS(tt.args.nums), "findLHS(%v)", tt.args.nums)
		})
	}
}

func Test_findLHS2(t *testing.T) {
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
				nums: []int{1, 4, 1, 3, 1, -14, 1, -13},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{1, 2, 1, 3, 0, 0, 2, 2, 1, 3, 3},
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{-3, -1, -1, -1, -3, -2},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 3, 5, 7, 9, 11, 13, 15, 17},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{1, 3, 2, 2, 5, 2, 3, 7},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findLHS2(tt.args.nums), "findLHS2(%v)", tt.args.nums)
		})
	}
}

func Test_findLHS3(t *testing.T) {
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
				nums: []int{1, 4, 1, 3, 1, -14, 1, -13},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{1, 2, 1, 3, 0, 0, 2, 2, 1, 3, 3},
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{-3, -1, -1, -1, -3, -2},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 3, 5, 7, 9, 11, 13, 15, 17},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{1, 3, 2, 2, 5, 2, 3, 7},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findLHS3(tt.args.nums), "findLHS3(%v)", tt.args.nums)
		})
	}
}
