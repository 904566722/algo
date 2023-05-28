package z3

import (
	"reflect"
	"testing"
)

func Test_differenceOfDistinctValues(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				grid: [][]int{
					{1, 2, 3},
					{3, 1, 5},
					{3, 2, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfDistinctValues(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("differenceOfDistinctValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_differenceOfDistinctValues2(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				grid: [][]int{
					{1, 2, 3},
					{3, 1, 5},
					{3, 2, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfDistinctValues2(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("differenceOfDistinctValues2() = %v, want %v", got, tt.want)
			}
		})
	}
}
