package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isMonotonic(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{1, 2, 2, 4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isMonotonic(tt.args.nums), "isMonotonic(%v)", tt.args.nums)
		})
	}
}
