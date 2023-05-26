package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				flowerbed: []int{0, 1, 0},
				n:         1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canPlaceFlowers(tt.args.flowerbed, tt.args.n), "canPlaceFlowers(%v, %v)", tt.args.flowerbed, tt.args.n)
		})
	}
}
