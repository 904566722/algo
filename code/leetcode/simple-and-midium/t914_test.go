package simple_and_midium

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasGroupsSizeX(t *testing.T) {
	type args struct {
		deck []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				deck: []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3},
			},
			want: false,
		}, {
			args: args{
				deck: []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2},
			},
			want: true,
		},
		{
			args: args{
				deck: []int{1, 1, 1, 2, 2, 2, 3, 3},
			},
			want: false,
		}, {
			args: args{
				deck: []int{1, 1, 2, 2, 2, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, hasGroupsSizeX(tt.args.deck), "hasGroupsSizeX(%v)", tt.args.deck)
		})
	}
}

func Test_zhengchu(t *testing.T) {
	fmt.Println(2 % 6)
	fmt.Println(6 % 2)
}

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: 4,
				b: 6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, gcd(tt.args.a, tt.args.b), "gcd(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
