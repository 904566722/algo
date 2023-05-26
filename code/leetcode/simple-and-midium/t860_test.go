package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				bills: []int{5, 5, 5, 10, 20},
			},
			want: true,
		}, {
			args: args{
				bills: []int{10, 5, 5, 5, 10, 20},
			},
			want: false,
		}, {
			args: args{
				bills: []int{5, 5, 5, 20, 10, 20},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lemonadeChange(tt.args.bills), "lemonadeChange(%v)", tt.args.bills)
		})
	}
}
