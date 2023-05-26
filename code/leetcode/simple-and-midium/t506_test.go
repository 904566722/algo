package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_award(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, award(tt.args.n), "award(%v)", tt.args.n)
		})
	}
}

func Test_findRelativeRanks(t *testing.T) {
	type args struct {
		score []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				score: []int{4, 5, 3, 2, 1},
			},
			want: []string{"Silver Medal", "Gold Medal", "Bronze Medal", "4", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findRelativeRanks(tt.args.score), "findRelativeRanks(%v)", tt.args.score)
		})
	}
}
