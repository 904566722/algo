package sort

import "testing"

func Test_quickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				nums: []int{5,2,34,56,7,8,2,23,4,6,7,8,10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums, 0, len(tt.args.nums) -1)
			t.Log(tt.args.nums)
		})
	}
}
