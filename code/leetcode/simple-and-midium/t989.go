package simple_and_midium

func addToArrayForm(num []int, k int) []int {
	idx := len(num) - 1
	var ans []int
	for k != 0 && idx >= 0 {
		lastNum := k%10 + num[idx]
		k = k / 10
		if lastNum >= 10 {
			k++
		}
		ans = append(ans, lastNum%10)
		idx--
	}

	for ; k != 0; k /= 10 {
		ans = append(ans, k%10)
	}
	for ; idx >= 0; idx-- {
		ans = append(ans, num[idx])
	}

	return reverse(ans)
}

func reverse(nums []int) []int {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
	return nums
}
