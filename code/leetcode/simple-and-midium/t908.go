package simple_and_midium

import (
	"leetcode/utils"
	"math"
)

func smallestRangeI(nums []int, k int) int {
	maxN, minN := math.MinInt, math.MaxInt
	for _, v := range nums {
		maxN = utils.Max(maxN, v)
		minN = min(minN, v)
	}
	h := maxN - minN
	if h <= k*2 {
		return 0
	} else {
		return maxN - minN - 2*k
	}
}
