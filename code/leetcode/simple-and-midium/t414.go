package simple_and_midium

import "sort"

func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	left, right := 0, 0 // left 左边代表已经去重的数组,right指针只管往右遍历，判断条件
	for ; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	// 倒数第三个数的下标
	tgtIdx := left - 2
	if tgtIdx < 0 {
		return nums[left]
	}
	return nums[tgtIdx]
}

// thirdMax2
// 用有点类似于 滑动窗口 的思想，一直维护一个只有三个值的数组或者 map 都可以
func thirdMax2(nums []int) int {
	// todo
	return 0
}
