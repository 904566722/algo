package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	type args struct {
		num []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				num: []int{0},
				k:   10000,
			},
			want: []int{1, 0, 0, 0},
		},
		{
			args: args{
				num: []int{2, 1, 5},
				k:   806,
			},
			want: []int{1, 0, 2, 1},
		},
		{
			args: args{
				num: []int{1, 2, 0, 0},
				k:   10,
			},
			want: []int{1, 2, 1, 0},
		}, {
			args: args{
				num: []int{1, 2},
				k:   10,
			},
			want: []int{2, 2},
		}, {
			args: args{
				num: []int{1, 2},
				k:   100,
			},
			want: []int{1, 1, 2},
		}, {
			args: args{
				num: []int{1, 2},
				k:   100,
			},
			want: []int{1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, addToArrayForm(tt.args.num, tt.args.k), "addToArrayForm(%v, %v)", tt.args.num, tt.args.k)
		})
	}
}
