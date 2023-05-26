package simple_and_midium

func minSubArrayLen(target int, nums []int) int {
	start, end := 0, 0
	n := len(nums)
	ans := n + 1
	sum := 0
	for end < n {
		for end < n && sum < target {
			sum += nums[end]
			end++
		}
		// [start, end)
		ans = min(ans, end-start)

		for start < end && sum >= target {
			sum -= nums[start]
			start++
		}
		ans = min(ans, end-start+1)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}