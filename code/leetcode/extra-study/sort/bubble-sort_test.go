package sort

import (
	"fmt"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				nums: []int{1,10,2,4,5,3,5,8,2,1},
			},
		},
	}
	bubbleSort(tests[0].args.nums)
	fmt.Println(tests[0].args.nums)
}
