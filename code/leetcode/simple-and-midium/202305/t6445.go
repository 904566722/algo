package _02305

import "leetcode/utils"

// if s[i] != s[i-1] 需要反转
//
//	#1. 反转 s[:i]
//	#2. 反转 s[i:]
//
// 选择较小的反转方式
// 无论是 #1 还是 #2 的反转方式，i 与 i+1 两个位置的字符是「一起变化」或者「不变化」
// 因此无需实际反转字符串
func minimumCost(s string) int64 {
	ans, n := int64(0), len(s)
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			ans += int64(utils.Min(i, n-i))
		}
	}
	return ans
}
