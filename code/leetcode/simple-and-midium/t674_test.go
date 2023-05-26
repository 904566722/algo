package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findLengthOfLCIS(t *testing.T) {
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
				nums: []int{1, 3, 5, 4, 7},
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1, 3, 5, 7},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findLengthOfLCIS(tt.args.nums), "findLengthOfLCIS(%v)", tt.args.nums)
		})
	}
}

func Test_findLengthOfLCIS2(t *testing.T) {
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
				nums: []int{1, 3, 5, 4, 7},
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1, 3, 5, 7},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findLengthOfLCIS2(tt.args.nums), "findLengthOfLCIS2(%v)", tt.args.nums)
		})
	}
}
