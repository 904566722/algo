package sort

func bubbleSort(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	// nums: 6 10 2 4 7 3 2 1 2
	// loop1                  j (loop2每进行一轮往前移，直到1)
	// loop2 i
	for j := len(nums)-1; j >= 1; j--{
		for i := 0; i < j; i++ {
			if nums[i] > nums[i+1] {
				swap(&nums[i], &nums[i+1])
			}
		}
	}
}