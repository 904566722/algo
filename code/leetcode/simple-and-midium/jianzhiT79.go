package simple_and_midium


func subsets(nums []int) [][]int {
	var ans [][]int
	for i := 0; i <= 1<<len(nums) - 1; i++ {
		var set []int
		for mask, j := i, 0; j < len(nums); j++ {
			if mask>>j&1 > 0 {
				set = append(set, nums[j])
			}
		}
		ans = append(ans, set)
	}
	return ans
}