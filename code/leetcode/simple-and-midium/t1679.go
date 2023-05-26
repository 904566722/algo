package simple_and_midium

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	cnt := 0
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum < k {
			i++
		} else if sum > k {
			j--
		} else {
			cnt++
			i++
			j--
		}
	}
	return cnt
}
