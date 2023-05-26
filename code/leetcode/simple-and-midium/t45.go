package simple_and_midium

func jump(nums []int) int {
	start := 0
	end := 1
	cnt := 0
	for end < len(nums) {
		maxPos := 0
		for i := start; i < end; i++ {
			maxPos = max(maxPos, i+nums[i])
		}
		start = end
		end = maxPos + 1
		cnt++
	}
	return cnt

}
