package simple_and_midium

import "sort"

// missingNumber 最简单的思路：下标标记
func missingNumber(nums []int) int {
	flg := make([]int, len(nums)+1)

	for _, v := range nums {
		flg[v] = 1
	}
	for i, v := range flg {
		if v == 0 {
			return i
		}
	}
	return -1
}

func missingNumber2(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 1 {
		return 1 - nums[0]
	}
	if nums[0] != 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			return nums[i] - 1
		}
	}
	return len(nums)
}

// missingNumber3 在这个数组后添加从 0 到 n 到数，总共 n + n + 1 = 2n + 1 个数
// 那么之前确实到数在这个新到数组中就只出现了一次，可以使用异或运算到特性来找到这个出现次数为奇数到数
func missingNumber3(nums []int) int {
	rst := 0
	for i := 0; i < len(nums); i++ {
		rst ^= nums[i]
		rst ^= i
	}
	rst ^= len(nums)
	return rst
}
