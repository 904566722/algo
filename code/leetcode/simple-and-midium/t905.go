package simple_and_midium

func sortArrayByParity(nums []int) []int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j]%2 == 0 {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	return nums
}
