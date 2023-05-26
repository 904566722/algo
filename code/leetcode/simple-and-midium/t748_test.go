package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rmDigitSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "1s3 PSt",
			},
			want: "psst",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, rmDigitSpaceLower(tt.args.s), "rmDigitSpace(%v)", tt.args.s)
		})
	}
}

func Test_shortestCompletingWord(t *testing.T) {
	type args struct {
		licensePlate string
		words        []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				licensePlate: "1s3 PSt",
				words:        []string{"step", "steps", "stripe", "stepple"},
			},
			want: "steps",
		}, {
			args: args{
				licensePlate: "1s3 456",
				words:        []string{"looks", "pest", "stew", "show"},
			},
			want: "pest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, shortestCompletingWord(tt.args.licensePlate, tt.args.words), "shortestCompletingWord(%v, %v)", tt.args.licensePlate, tt.args.words)
		})
	}
}
