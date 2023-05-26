package simple_and_midium

func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != 0 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] != 0 {
				nums[i] = nums[j]
				nums[j] = 0
				break
			}
		}
	}
}

// moveZeroes2 双指针
// 在同一个数组中做操作，将数组中到数分为两类：零、非零，分别使用两个指针来表示
func moveZeroes2(nums []int) {
	left, right := 0, 0
	for ; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}
	for ; left < len(nums); left++ {
		nums[left] = 0
	}
}
