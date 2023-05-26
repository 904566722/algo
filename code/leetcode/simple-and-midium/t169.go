package simple_and_midium

import (
	"math"
	"sort"
)

// 本质上是统计数组中数字出现的次数
// 找到数组中出现次数多于半数(n/2)的数
// 使用 map 来存储，k-v，k:nums[i] v：用来计数
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	rst := math.MinInt
	cnt := make(map[int]int)
	for _, num := range nums {
		if _, ok := cnt[num]; ok {
			cnt[num] += 1

			if cnt[num] > len(nums)/2 {
				rst = num
			}
		} else {
			cnt[num] = 1
		}
	}

	return rst
}

// majorityElement2 题目设定必定有多数元素，那么当数组排序之后，多数元素必定在数组中间
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
