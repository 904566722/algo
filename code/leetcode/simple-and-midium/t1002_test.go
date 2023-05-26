package simple_and_midium

import (
	"reflect"
	"testing"
)

func Test_rmDuplication(t *testing.T) {
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
				s: "aabccdd",
			},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rmDuplication(tt.args.s); got != tt.want {
				t.Errorf("rmDuplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commonChars(t *testing.T) {
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
				words: []string{"bella", "label", "roller"},
			},
			want: []string{"e", "l", "l"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_codingP(t *testing.T) {
	codingP()
}
