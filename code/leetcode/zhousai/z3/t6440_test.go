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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfDistinctValues(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("differenceOfDistinctValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
