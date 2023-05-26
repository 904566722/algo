package simple_and_midium

import (
	"leetcode/utils"
	"math"
)

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	ans := math.MinInt
	for l1 := 0; l1+firstLen-1 < len(nums); l1++ {
		sum1 := 0
		r1 := l1 + firstLen - 1
		for _, v := range nums[l1 : l1+firstLen] {
			sum1 += v
		}

		for l2 := 0; l2+secondLen-1 < len(nums); l2++ {
			r2 := l2 + secondLen - 1
			sum2 := 0
			// 存在交集的情况 以及 右边的长度已经不够，跳过
			if (l1 <= l2 && l2 <= r1) || (l1 <= r2 && r2 <= r1) || (l2 < l1 && r2 > r1) {
				continue
			}

			for _, v := range nums[l2 : r2+1] {
				sum2 += v
			}
			if r2+1 > len(nums) {
				continue
			}
			ans = utils.Max(ans, sum1+sum2)
		}
	}
	return ans
}
