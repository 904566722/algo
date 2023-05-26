package simple_and_midium

import "strings"

// 字典中的字符串按长度从大到小排列
//	假设：aaaopen,open，aaaa
//	如果字符串中有 aaaopen，一定要先用 aaaopen 来构成，而不是 open + aaa
//	s：aaaaopen，裁剪 aaaopen 之后剩下 a，但是 s 是可以由 aaaa + open 组成的
//	所以上面的思路行不通！
//
// 如果用递归的思路来做
//	按照给定的字典，把所有可能的情况都裁剪一遍
//	如果最后 s 为 “” true
//	如果没有可以裁剪的子字符串，false
//	超时！
//
// 如果列举出所有 wordDict 能够组成的单词，再来比较怎么样？ n <= 1000
// 比递归还慢吧
//
//						i-len(word)+1
// s： x	x	x	x	x	x	x	x	x	x	x	x
//								i	想要判断 s[:i+1] 是否合法，遍历字典，
//						|	-	|
//									如果存在一个 word，当 s[:i+1] - word 剩下的字符串是合法的，那么 s[:i+1] 就合法
// 使用 dp[i] 来保存 s[:i+1] 是否合法
//	success
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		for _, word := range wordDict {
			m := len(word)
			if i-m+1 < 0 {
				continue
			}
			if s[i-m+1:i+1] == word && i-m+1 == 0 {
				// 考虑第一个字母为一个单词的情况
				dp[i] = 1
				break
			} else if s[i-m+1:i+1] == word && dp[i-m] == 1 {
				dp[i] = 1
				break
			}
		}

	}
	return dp[n-1] == 1
}

// 递归
func cut(s string, wordDict []string) bool {
	if s == "" {
		return true
	}

	for _, word := range wordDict {
		idx := strings.Index(s, word)
		if idx >= 0 {
			if cut(s[:idx], wordDict) && cut(s[idx+len(word):], wordDict) {
				return true
			}
		}
	}
	return false
}