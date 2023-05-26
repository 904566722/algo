package z1

import (
	"testing"
)

func Test_firstCompleteIndex(t *testing.T) {
	type args struct {
		arr []int
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr: []int{1,3,4,2},
				mat: [][]int{
					{1,4},
					{2,3},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstCompleteIndex(tt.args.arr, tt.args.mat); got != tt.want {
				t.Errorf("firstCompleteIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstCompleteIndex2(t *testing.T) {
	type args struct {
		arr []int
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr: []int{8,2,4,9,3,5,7,10,1,6},
				mat: [][]int{
					{8,2,9,10,4},
					{1,7,6,3,5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstCompleteIndex2(tt.args.arr, tt.args.mat); got != tt.want {
				t.Errorf("firstCompleteIndex2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstCompleteIndex3(t *testing.T) {
	type args struct {
		arr []int
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr: []int{10,12,1,7,2,6,9,11,8,5,3,4},
				mat: [][]int{
					{8,1,6,10},
					{5,9,2,4},
					{12,3,7,11},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstCompleteIndex3(tt.args.arr, tt.args.mat); got != tt.want {
				t.Errorf("firstCompleteIndex3() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_compEfficiency(t *testing.T) {
	compEfficiency()
}

func Test_testBreak(t *testing.T) {
	testBreak()
}