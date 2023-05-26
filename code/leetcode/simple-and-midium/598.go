package simple_and_midium

import "math"

func maxCount(m int, n int, ops [][]int) int {
	l := len(ops)
	if l == 0 {
		return m * n
	}

	minR, minC := math.MaxInt, math.MaxInt
	for _, op := range ops {
		if minR > op[0] {
			minR = op[0]
		}
		if minC > op[1] {
			minC = op[1]
		}
	}

	return minR * minC
}
