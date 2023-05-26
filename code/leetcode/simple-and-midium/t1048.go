package simple_and_midium

import (
	"leetcode/utils"
	"sort"
)

// 排序 O(nlogn)
// 获取 idx 列表 O(n)
// 遍历 O(n) * O(l) 每个单词比较
func longestStrChain(words []string) int {
	if len(words) == 1 {
		return 1
	}
	sortWords(words)

	// idxList[l] 保存了单词长度为 l 的下标列表
	// 长度为最大单词长度 + 1
	idxList := getIdxList(words)
	n, ans := len(words), 1
	// dp[i] 代表以当前单词（words[i]）结尾的最长单词链的长度
	dp := make([]int, n)
	// cur 指向当前单词，pre 指向长度 -1 的单词
	for cur := 0; cur < n; cur++ {
		l := len(words[cur]) // 当前单词的长度
		preList := idxList[l-1]
		// 没有长度为 l - 1 的单词
		if len(preList) == 0 {
			dp[cur] = 1
		} else {
			// 遍历长度为 l-1 的单词，更新 dp
			for _, pre := range preList {
				if isQianshen(words[pre], words[cur]) {
					dp[cur] = utils.Max(dp[cur], dp[pre]+1)
				} else {
					dp[cur] = utils.Max(dp[cur], 1)
				}
				ans = utils.Max(ans, dp[cur])
			}
		}
	}

	return ans
}

// 按照单词的长度升序排序
func sortWords(words []string) {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
}

// 前提：words 按照单词长度排序
func getIdxList(words []string) [][]int {
	idxList := make([][]int, len(words[len(words)-1])+1)
	idxList[0] = []int{}
	for i, w := range words {
		idxList[len(w)] = append(idxList[len(w)], i)
	}
	return idxList
}

func isQianshen(short, long string) bool {
	n1, n2 := len(short), len(long)
	if n1 >= n2 || n1 != n2-1 {
		return false
	}

	jumpFlg := false
	for i, j := 0, 0; i < n1 && j < n2; {
		if jumpFlg && short[i] != long[j] {
			return false
		}
		if short[i] == long[j] {
			i++
			j++
		} else {
			if !jumpFlg {
				j++
				jumpFlg = true
			}
		}
	}
	return true
}
