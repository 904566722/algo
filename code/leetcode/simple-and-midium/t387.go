package simple_and_midium

func firstUniqChar(s string) int {
	cnt := make([]int, 26)
	for i := len(s) - 1; i >= 0; i-- {
		cnt[s[i]-'a']++
	}

	for i := range s {
		if cnt[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
