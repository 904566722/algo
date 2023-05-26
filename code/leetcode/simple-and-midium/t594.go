package simple_and_midium

import (
	"sort"
)

func findLHS(nums []int) int {
	n := len(nums)
	numsTmp := make([]int, n)
	copy(numsTmp, nums)

	sort.Ints(numsTmp)
	if n == 1 || numsTmp[0]-numsTmp[n-1] == 0 {
		return 0
	}

	maxLen := 0
	trueMaxLen := 0
	lenMp := make(map[int]int) // 以数字k为起点的最长和谐子序列的长度
	for i := 0; i < n; i++ {
		if _, ok := lenMp[numsTmp[i]]; ok {
			continue
		}
		maxLen = 0
		for j := i + 1; j < n && (numsTmp[j] == numsTmp[i]+1 || numsTmp[j] == numsTmp[i]); j++ {
			maxLen++
			lenMp[numsTmp[i]] = maxLen
			if trueMaxLen < maxLen && numsTmp[j] == numsTmp[i]+1 {
				trueMaxLen = maxLen
			}
		}
	}

	if trueMaxLen == 0 {
		return 0
	}
	return trueMaxLen + 1
}

// 滑动窗口
// 窗口起点移动条件，终点与起点差值大于1（不符合条件）
// 判断长度的时候起点与终点差值必须为1
func findLHS2(nums []int) int {
	sort.Ints(nums)
	left := 0
	maxLen := 0
	for right, n := range nums {
		for n > nums[left]+1 {
			left++
		}
		if nums[right] == nums[left]+1 && right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// 下标计数法
// 把题目看成计数的题目，也就很自然就想到了这个方法
func findLHS3(nums []int) int {
	mp := make(map[int]int)
	for _, v := range nums {
		if _, ok := mp[v]; ok {
			mp[v]++
		} else {
			mp[v] = 1
		}
	}

	maxLen := 0
	for k, v := range mp {
		if _, ok := mp[k+1]; ok && maxLen < v+mp[k+1] {
			maxLen = v + mp[k+1]
		}
	}

	return maxLen
}
