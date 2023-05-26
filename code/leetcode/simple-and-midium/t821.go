package simple_and_midium

import "math"

func shortestToChar(s string, c byte) []int {
	ans := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		ans[i] = math.MaxInt
	}
	if s[0] == c {
		ans[0] = 0
	}
	if s[len(s)-1] == c {
		ans[len(s)-1] = 0
	}

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			ans[i] = 0
		} else if i-1 >= 0 && ans[i-1] != math.MaxInt {
			ans[i] = min(ans[i], ans[i-1]+1)
		}

		right := len(s) - i - 1
		if s[right] == c {
			ans[right] = 0
		} else if right+1 < len(s) && ans[right+1] != math.MaxInt {
			ans[right] = min(ans[right], ans[right+1]+1)
		}
	}
	return ans
}
