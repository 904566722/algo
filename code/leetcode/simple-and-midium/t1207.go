package simple_and_midium

func uniqueOccurrences(arr []int) bool {
	freq := map[int]int{}
	for _, v := range arr {
		freq[v]++
	}
	freqCnt := make([]int, 1001)
	for _, v := range freq {
		if freqCnt[v] > 0 {
			return false
		}
		freqCnt[v]++
	}
	return true
}