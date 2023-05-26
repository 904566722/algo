package simple_and_midium

func convertToTitle(columnNumber int) string {
	letters := make([]byte, 26)
	for i := 0; i < len(letters); i++ {
		letters[i] = byte(i) + 'A'
	}

	var ans []byte
	for columnNumber > 0 {
		columnNumber--
		ans = append(ans, letters[columnNumber%26])
		columnNumber /= 26
	}

	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return string(ans)
}
