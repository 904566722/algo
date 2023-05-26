package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			},
			want: 6,
		}, {
			args: args{
				cost: []int{10, 15, 20},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minCostClimbingStairs(tt.args.cost), "minCostClimbingStairs(%v)", tt.args.cost)
		})
	}
}
