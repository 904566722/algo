package simple_and_midium

func rotate(nums []int, k int)  {
	n := len(nums)
	if k == 0 || k == n {
		return
	}
	k = k % n

	cacheNum := nums[0]
	for i := 0; i < k; i++ {
		nextId := i + k
		for nextId < n {
			nums[nextId], cacheNum = cacheNum, nums[nextId]
			nextId += k
		}
		nextId %= n
		nums[nextId], cacheNum = cacheNum, nums[nextId]
	}
}
