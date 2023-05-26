package sort

func sort(nums []int)  {
	quickSort(nums, 0, len(nums) - 1)
}

// nums 数组排序
func quickSort(nums []int, start, end int) {
	// 结束条件
	if start >= end {
		return
	}

	// 子问题关系
	pivot := partition(nums, start, end)
	quickSort(nums, start, pivot)
	quickSort(nums, pivot+1, end)
	return
}

func partition(nums []int, start, end int) int {
	if start == end {

	}
	pivot := start
	idx := start + 1
	for i := start+1; i <= end; i++ {
		if nums[i] < nums[pivot] {
			swap(&nums[i], &nums[idx])
			idx++
		}
	}
	swap(&nums[pivot], &nums[idx-1])
	return idx-1
}
