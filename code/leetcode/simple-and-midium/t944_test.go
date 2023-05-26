package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minDeletionSize(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				strs: []string{"cba", "daf", "ghi"},
			},
			want: 1,
		}, {
			args: args{
				strs: []string{"a", "b"},
			},
			want: 0,
		}, {
			args: args{
				strs: []string{"zyx", "wvu", "tsr"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minDeletionSize(tt.args.strs), "minDeletionSize(%v)", tt.args.strs)
		})
	}
}
