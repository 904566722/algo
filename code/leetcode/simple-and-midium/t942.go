package simple_and_midium

func diStringMatch(s string) []int {
	n := len(s)
	ans := make([]int, n+1)
	nums := make([]int, n+1)
	for i := 0; i <= n; i++ {
		nums[i] = i
	}
	for i, ch := range s {
		if ch == 'I' {
			ans[i] = nums[0]
			nums = nums[1:]
		} else {
			ans[i] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		}
	}
	ans[n] = nums[0]
	return ans
}
