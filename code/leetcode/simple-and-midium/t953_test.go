package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				words: []string{"abc"},
				order: "ab",
			},
		}, {
			args: args{
				words: []string{"abc"},
				order: "你好",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isAlienSorted(tt.args.words, tt.args.order), "isAlienSorted(%v, %v)", tt.args.words, tt.args.order)
		})
	}
}
