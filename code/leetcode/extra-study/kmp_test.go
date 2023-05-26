package extra_study

import (
	"reflect"
	"testing"
)

func Test_getNext(t *testing.T) {
	type args struct {
		pat string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				pat: "a",
			},
			want: []int{0},
		},
		{
			args: args{
				pat: "aabaac",
			},
			want: []int{0, 1, 0, 1, 2, 0},
		}, {
			args: args{
				pat: "ababc",
			},
			want: []int{0, 0, 1, 2, 0},
		}, {
			args: args{
				pat: "afdabeafdaf",
			},
			want: []int{0, 0, 0, 1, 0, 0, 1, 2, 3, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNext(tt.args.pat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMP(t *testing.T) {
	type args struct {
		s   string
		pat string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s:   "abab",
				pat: "ba",
			},
			want: 1,
		},
		{
			args: args{
				s:   "abab",
				pat: "ab",
			},
			want: 0,
		},
		{
			args: args{
				s:   "aabaabaac",
				pat: "aabaac",
			},
			want: 3,
		}, {
			args: args{
				s:   "abab",
				pat: "ab",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KMP(tt.args.s, tt.args.pat); got != tt.want {
				t.Errorf("KMP() = %v, want %v", got, tt.want)
			}
		})
	}
}
