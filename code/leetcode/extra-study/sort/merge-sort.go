package sort

// 1. 定义方法：将左右两个已经有序的序列合并成一个序列（原问题得解：原数组有序）
// 2. 子问题：左右的数组需要有序，与原问题解决方法相同
// 3. 结束条件：数组的长度为 1
func mergeSort(nums []int) []int {
	// 结束条件
	if len(nums) <= 1 {
		return nums
	}

	// 子问题关系
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	// 方法定义
	return merge(left, right)
}

// 将两个有序数组合并
func merge(nums1, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	nums := make([]int, n1 + n2)
	i, j, idx := 0, 0, 0
	for i < n1 && j < n2 {
		if nums1[i] < nums2[j] {
			nums[idx] = nums1[i]
			i++
		} else {
			nums[idx] = nums2[j]
			j++
		}
		idx++
	}
	for ; i < n1; i++ {
		nums[idx] = nums1[i]
		idx++
	}
	for ; j < n2; j++ {
		nums[idx] = nums2[j]
		idx++
	}

	return nums
}


// 递归 practice
func jump(n int) int {
	if n <= 2 {
		return n
	}
	return jump(n-1) + jump(n-2)
}