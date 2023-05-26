package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_smallestRangeI(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1},
				k:    0,
			},
			want: 0,
		}, {
			args: args{
				nums: []int{0, 10},
				k:    2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, smallestRangeI(tt.args.nums, tt.args.k), "smallestRangeI(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
