package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dominantIndex(t *testing.T) {
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
				nums: []int{3, 6, 1, 0},
			},
			want: 1,
		}, {
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: -1,
		}, {
			args: args{
				nums: []int{1},
			},
			want: 0,
		}, {
			args: args{
				nums: []int{0, 0, 3, 2},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, dominantIndex(tt.args.nums), "dominantIndex(%v)", tt.args.nums)
		})
	}
}
