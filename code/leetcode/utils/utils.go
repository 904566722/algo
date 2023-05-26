package utils

import "math"

func Max(nums ...int) int {
	rst := math.MinInt
	for _, v := range nums {
		if rst < v {
			rst = v
		}
	}
	return rst
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
