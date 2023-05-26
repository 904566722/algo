package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calPoints(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				operations: []string{"5", "2", "C", "D", "+"},
			},
			want: 30,
		},
		{
			args: args{
				operations: []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, calPoints(tt.args.operations), "calPoints(%v)", tt.args.operations)
		})
	}
}
