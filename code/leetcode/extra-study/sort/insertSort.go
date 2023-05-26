package sort

func insertSort(nums []int)  {
	// nums 1	2	5	3	4
	//          i (i极其右边为一个 未排序的序列)
	//      j=i-1 (j 负责往前找大于等于cur的值)
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		j := i-1
		for ; j>=0 && nums[j]>cur; j-- {
			nums[j+1] = nums[j]
		}
		// 由于 j 当前的位置小于或者等于cur or 越界
		// 故把前面一个位置，给到cur（前面的数已经往前前的位置移动）
		nums[j+1] = cur
	}
}