package simple_and_midium

import "strconv"

func freqAlphabets(s string) string {
	n2letter := make([]byte, 26)
	for i := byte(0); i<26; i++ {
		n2letter[i] = i + 'a'
	}

	var ans []byte
	for i := len(s) - 1; i>=0; {
		var num int
		if s[i] == '#' {
			num, _ = strconv.Atoi(s[i-2 : i])
			i -= 3
		} else {
			num, _ = strconv.Atoi(string(s[i]))
			i--
		}
		ans = append(ans, n2letter[num-1])
	}
	for i, n := 0, len(ans); i < n / 2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return string(ans)
}
