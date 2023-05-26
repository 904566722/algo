package simple_and_midium

import (
	"leetcode/utils"
	"math"
)

// 等差子序列，最长的长度
// f[nums[i]] 表示以 nums[i] 为结尾的等差子序列的最长长度
// f[nums[i]] = f[nums[i]-d] + 1
// 			 or = 1
func longestArithSeqLength(nums []int) int {
	maxN, minN := math.MinInt, math.MaxInt
	for _, n := range nums {
		if n > maxN {
			maxN = n
		}
		if n < minN {
			minN = n
		}
	}

	dRange := maxN - minN
	rst := math.MinInt
	// 遍历所有的差值情况
	for d := -dRange; d <= dRange; d++ {
		f := make([]int, maxN+1)
		for _, num := range nums {
			pre := num - d
			if pre >= minN && pre <= maxN && f[pre] > 0 {
				f[num] = utils.Max(f[pre]+1, f[num])
			}
			f[num] = utils.Max(1, f[num])
			rst = utils.Max(f[num], rst)
		}
	}
	return rst
}
