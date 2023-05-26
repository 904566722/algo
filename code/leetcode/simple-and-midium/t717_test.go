package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isOneBitCharacter(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				bits: []int{1, 0, 0},
			},
			want: true,
		}, {
			args: args{
				bits: []int{1, 1, 1, 0},
			},
			want: false,
		}, {
			args: args{
				bits: []int{1, 1},
			},
			want: false,
		}, {
			args: args{
				bits: []int{1, 1, 0, 0},
			},
			want: true,
		}, {
			args: args{
				bits: []int{0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isOneBitCharacter(tt.args.bits), "isOneBitCharacter(%v)", tt.args.bits)
		})
	}
}
