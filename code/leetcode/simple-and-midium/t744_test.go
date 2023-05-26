package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		//{
		//	args: args{
		//		letters: []byte{'c', 'f', 'j'},
		//		target:  'a',
		//	},
		//	want: 'c',
		//},
		{
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'c',
			},
			want: 'f',
		}, {
			args: args{
				letters: []byte{'x', 'x', 'z', 'z'},
				target:  'z',
			},
			want: 'x',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, nextGreatestLetter(tt.args.letters, tt.args.target), "nextGreatestLetter(%v, %v)", tt.args.letters, tt.args.target)
		})
	}
}

func Test_nextGreatestLetter2(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'c',
			},
			want: 'f',
		}, {
			args: args{
				letters: []byte{'x', 'x', 'z', 'z'},
				target:  'z',
			},
			want: 'x',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, nextGreatestLetter2(tt.args.letters, tt.args.target), "nextGreatestLetter2(%v, %v)", tt.args.letters, tt.args.target)
		})
	}
}
