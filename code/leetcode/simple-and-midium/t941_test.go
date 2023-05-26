package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				arr: []int{3, 5, 5},
			},
			want: false,
		}, {
			args: args{
				arr: []int{3, 5, 1},
			},
			want: true,
		}, {
			args: args{
				arr: []int{3, 5, 5},
			},
			want: false,
		}, {
			args: args{
				arr: []int{3, 5, 5},
			},
			want: false,
		}, {
			args: args{
				arr: []int{3, 3, 5},
			},
			want: false,
		}, {
			args: args{
				arr: []int{1, 2, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, validMountainArray(tt.args.arr), "validMountainArray(%v)", tt.args.arr)
		})
	}
}
