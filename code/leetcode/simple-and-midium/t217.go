package simple_and_midium

import "sort"

// containsDuplicate 先排序，然后判断每个相邻的数是否存在相同的
// 时间主要花费在排序的时间上，剩下的只需要遍历一边数组
// 空间复杂度:o(1)
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	rst := false
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			rst = true
			break
		}
	}
	return rst
}

// containsDuplicate2
// 这题也可以用 map 来实现，省去解法1中的排序步骤，把这个题目看成单纯的计数题目
func containsDuplicate2(nums []int) bool {
	rst := false
	cnt := make(map[int]int)
	for _, num := range nums {
		if _, ok := cnt[num]; ok {
			rst = true
			break
		} else {
			cnt[num] = 1
		}
	}

	return rst
}
