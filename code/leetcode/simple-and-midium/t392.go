package simple_and_midium

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	i := 0
	for j := 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

// dp
func isSubsequence2(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}
	dp := make([][26]int, len(t)+1)
	for i := 0; i < 26; i++ {
		dp[len(t)][i] = -1
	}
	for i := len(t) - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			ch := byte(j) + 'a'
			if t[i] == ch {
				dp[i][j] = i
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	i, j := 0, dp[0][s[0]-'a']
	for i < len(s) && j < len(t) && 0 <= j {
		if s[i] == t[j] {
			i++
		}
		if i >= len(s) {
			return true
		}
		j = dp[j+1][s[i]-'a']
	}
	return i >= len(s)
}
