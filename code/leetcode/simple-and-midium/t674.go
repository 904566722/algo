package simple_and_midium

import (
	"leetcode/utils"
	"math"
)

func findLengthOfLCIS(nums []int) int {
	//if len(nums) == 1 {
	//	return 1
	//}
	start, end, pre := 0, 0, nums[0]
	maxLen := math.MinInt
	for ; end < len(nums); end++ {
		// 左端点移动的条件，end对应的数值不变或者突然下降
		if nums[end] <= pre {
			if maxLen < end-start {
				maxLen = end - start
			}
			start = end
		}
		pre = nums[end]
	}
	if maxLen < end-start {
		maxLen = end - start
	}
	return maxLen
}

func findLengthOfLCIS2(nums []int) int {
	start := 0
	maxLen := 0
	for i, v := range nums {
		if i > 0 && v <= nums[i-1] {
			start = i
		}
		maxLen = utils.Max(maxLen, i-start+1)
	}
	return maxLen
}
