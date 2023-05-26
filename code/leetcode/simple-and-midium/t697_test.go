package simple_and_midium

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_findShortestSubArray(t *testing.T) {
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
				nums: []int{1, 2, 2, 3, 1},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{1, 2, 2, 3, 1, 4, 2},
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{1},
			},
			want: 1,
		}, {
			args: args{
				nums: []int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findShortestSubArray(tt.args.nums), "findShortestSubArray(%v)", tt.args.nums)
		})
	}
}

func TestC(t *testing.T) {
	fmt.Println(math.Pow10(6) + 1)
}
