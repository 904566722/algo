package simple_and_midium

//
//func maxScore(s string) int {
//	n := len(s)
//	zeros, ones := make([]int, n), make([]int, n)
//	if s[0] == '0' {
//		zeros[0] = 1
//	}
//
//	for i := 1; i < n; i++ {
//		addOne := 0
//		if s[i] == '0' {
//			addOne = 1
//		}
//		zeros[i] = zeros[i-1] + 1 * addOne
//	}
//	for j := n - 2; j >= 0; j-- {
//		addOne := 0
//		if s[j + 1] == '1' {
//			addOne = 1
//		}
//		ones[j] = ones[j+1] + 1 * addOne
//	}
//
//	ans := math.MinInt
//	for i := 0; i < n - 1; i++ {
//		ans = string2.max(ans, zeros[i] + ones[i])
//	}
//	return ans
//}
