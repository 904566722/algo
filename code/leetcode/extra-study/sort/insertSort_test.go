package sort

import "testing"

func Test_insertSort(t *testing.T) {
	insertSort([]int{3, 44, 38, 15, 2, 5, 3})
}

func Test_insertSort1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}