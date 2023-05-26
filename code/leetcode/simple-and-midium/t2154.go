package simple_and_midium

import "math"

func findFinalValue(nums []int, original int) int {
	appeared := map[int]bool{}
	maxNum := math.MinInt
	for _, num := range nums {
		appeared[num] = true
		if maxNum < num {
			maxNum = num
		}
	}
	for original <= maxNum && appeared[original] {
		original *= 2
	}
	return original
}