package simple_and_midium

func repeatedNTimes(nums []int) int {
	cnt := make(map[int]bool)
	for _, n := range nums {
		if _, ok := cnt[n]; ok {
			return n
		} else {
			cnt[n] = true
		}
	}
	return -1
}
