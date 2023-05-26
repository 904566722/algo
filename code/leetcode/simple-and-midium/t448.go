package simple_and_midium

import "sort"

// 下标标记法
func findDisappearedNumbers(nums []int) []int {
	var rst []int
	flg := make([]int, len(nums)+1)
	for _, v := range nums {
		flg[v] = 1
	}
	for i := 1; i < len(flg); i++ {
		if flg[i] == 0 {
			rst = append(rst, i)
		}
	}
	return rst
}

// time: o(n)
func findDisappearedNumbers2(nums []int) []int {
	var rst []int
	sort.Ints(nums)
	for j := 1; j < nums[0]; j++ {
		rst = append(rst, j)
	}
	for i := 0; i < len(nums)-1; i++ {
		for missNum := nums[i]; missNum+1 != nums[i+1] && missNum != nums[i+1]; missNum++ {
			rst = append(rst, missNum+1)
		}
	}
	for j := nums[len(nums)-1] + 1; j <= len(nums); j++ {
		rst = append(rst, j)
	}
	return rst
}
