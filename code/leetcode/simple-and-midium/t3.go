package simple_and_midium

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	ans := 0
	for i := 0; i < n; i++ {
		freq := make(map[byte]int)
		freq[s[i]]++
		j := i + 1
		for ; j < n; j++ {
			freq[s[j]]++
			if freq[s[j]] >= 2 {
				ans = max(ans, j-i)
				break
			}
		}
		if j == n {
			ans = max(ans, j-i)
		}
	}
	return ans
}
