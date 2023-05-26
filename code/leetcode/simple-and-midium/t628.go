package simple_and_midium

import "sort"

func maximumProduct(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	if nums[n-1] >= 0 {
		if nums[n-2]*nums[n-3] > nums[0]*nums[1] {
			return nums[n-1] * nums[n-2] * nums[n-3]
		}
		return nums[n-1] * nums[0] * nums[1]
	} else {
		return nums[n-1] * nums[n-2] * nums[n-3]
	}
}
