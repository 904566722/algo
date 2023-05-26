package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{2, 3, 3, 4, 5, 6},
			},
			want: []int{3, 1},
		},
		{
			args: args{
				nums: []int{1, 1},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findErrorNums(tt.args.nums), "findErrorNums(%v)", tt.args.nums)
		})
	}
}

func Test_findErrorNums2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{2, 3, 3, 4, 5, 6},
			},
			want: []int{3, 1},
		},
		{
			args: args{
				nums: []int{1, 1},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findErrorNums2(tt.args.nums), "findErrorNums2(%v)", tt.args.nums)
		})
	}
}
