package simple_and_midium

import "math"

// 1 找到数组的 度（元素出现频数的最大值）
//		可能有多个相同的度
// 2 找到一个子数组，拥有相同的度，且最短
func findShortestSubArray(nums []int) int {
	cnt := make(map[int]int)
	maxCnt := math.MinInt
	for _, v := range nums {
		cnt[v]++
		if maxCnt < cnt[v] {
			maxCnt = cnt[v]
		}
	}

	l := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if cnt[nums[i]] == maxCnt {
			j := len(nums) - 1
			for ; j > i && nums[j] != nums[i]; j-- {
			}
			if l > j-i+1 {
				l = j - i + 1
			}
			cnt[nums[i]]--
		}
	}

	return l
}
