package simple_and_midium

func findErrorNums(nums []int) []int {
	n := len(nums)
	mp := make(map[int]int)
	for i, _ := range nums {
		mp[i+1] = -1
	}
	rst := make([]int, 2)
	expectSum := n * (n + 1) / 2
	actualSum := 0
	for i := 0; i < len(nums); i++ {
		actualSum += nums[i]
		if mp[nums[i]] == -1 {
			mp[nums[i]]++
		} else {
			rst[0] = nums[i]
		}
	}
	rst[1] = expectSum - actualSum + rst[0]

	return rst
}

// 如果是用下标计数，不用上面那么复杂。。
func findErrorNums2(nums []int) []int {
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}

	rst := make([]int, 2)
	for i := 1; i <= len(nums); i++ {
		if v, ok := cnt[i]; ok && v == 2 {
			rst[0] = i
		} else if !ok {
			rst[1] = i
		}
	}

	return rst
}
