package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_imageSmoother(t *testing.T) {
	type args struct {
		img [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				img: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			},
			want: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			args: args{
				img: [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}},
			},
			want: [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, imageSmoother(tt.args.img), "imageSmoother(%v)", tt.args.img)
		})
	}
}
