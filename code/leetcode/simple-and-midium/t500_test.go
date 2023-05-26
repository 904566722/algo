package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_belongLine(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				b: "q",
			},
			want: 1,
		},
		{
			args: args{
				b: "g",
			},
			want: 2,
		},
		{
			args: args{
				b: "M",
			},
			want: 3,
		},
		{
			args: args{
				b: "z",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, belongLine(tt.args.b), "belongLine(%v)", tt.args.b)
		})
	}
}

func Test_findWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				words: []string{"Hello", "Alaska", "Dad", "Peace"},
			},
			want: []string{"Alaska", "Dad"},
		},
		{
			args: args{
				words: []string{"a", "b"},
			},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findWords(tt.args.words), "findWords(%v)", tt.args.words)
		})
	}
}
