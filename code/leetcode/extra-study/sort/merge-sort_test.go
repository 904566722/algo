package sort

import (
	"fmt"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{5,4,2,3,6,8,2,2,5,7,9},
			},
		},
	}
	fmt.Println(mergeSort(tests[0].args.nums))
}
