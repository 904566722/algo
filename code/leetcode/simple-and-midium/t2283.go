package simple_and_midium

import "strconv"

func digitCount(num string) bool {
	n := len(num)
	freq := make([]int, 10)
	sum := 0
	for i, ch := range num {
		n, _ := strconv.Atoi(string(ch))
		freq[i] = n
		sum += n
	}
	if sum != n {
		return false
	}
	for _, ch := range num {
		if freq[ch-'0']--; freq[ch-'0'] < 0 {
			return false
		}
	}
	return true
}