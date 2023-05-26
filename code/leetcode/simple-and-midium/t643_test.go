package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				nums: []int{1, 12, -5, -6, 50, 3},
				k:    4,
			},
			want: 12.75,
		},
		{
			args: args{
				nums: []int{5},
				k:    1,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMaxAverage(tt.args.nums, tt.args.k), "findMaxAverage(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
