package simple_and_midium

func removeDuplicates4(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	for i := 2; i < len(nums); {
		if nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			copy(nums[i:], nums[i+1:])
			nums = nums[:len(nums)-1]
		} else {
			i++
		}
	}
	return len(nums)
}
