package simple_and_midium

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i, v := range nums {
		if i%2 == 0 {
			sum += v
		}
	}
	return sum
}
