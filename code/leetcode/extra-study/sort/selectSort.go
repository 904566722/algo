package sort

func selectSort(nums []int)  {
	for i := 0; i < len(nums)-1; i++ {
		minIdx := i
		for j := i+1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				minIdx = j
			}
		}
		swap(&nums[i], &nums[minIdx])
	}
}
