package simple_and_midium

import (
	"leetcode/utils"
	"math"
)

// 设 dp[i] 表示 放入第 i 本书需要第最小高度
// 那么 dp[i+1] = min(dp[i], dp[i]+h)
//
// 如果把每一层书架看作是 books 第一个子集，问题转化为：求 books 第所有子集划分情况中，子集中第最大值的和 的 最小值
// 假设 books[i+1] 为目前子集的最后一个数，遍历这个子集可能的所有情况(结束的边界：w > shelfWidth or j < 0)
// 假设这个子集的起始元素 books[j]（0<=j<=i），这个子集为 books[j...i+1]，记录这个子集的最大高度 h
// 那么 dp[i+1] = min(dp[i+1], dp[j-1] + h)

func minHeightShelves(books [][]int, shelfWidth int) int {
	dp := make([]int, len(books)+1)
	for i := 0; i < len(books); i++ {
		dp[i] = math.MaxInt
		w, h := 0, 0

		// 遍历子集的可能情况，找到最小值 dp[i]
		for j := i; j >= 0; j-- {
			w += books[j][0]
			h = utils.Max(h, books[j][1])
			if w > shelfWidth {
				break
			}
			if j-1 < 0 {
				dp[i] = utils.Max(dp[i], h)
			} else {
				dp[i] = utils.Max(dp[i], dp[j-1]+h)
			}

		}
	}
	return dp[len(books)-1]
}
