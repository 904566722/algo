package simple_and_midium

// 字符串形式的两个数相加
// 1 1
// 1 1
func addBinary(a string, b string) string {
	var ans []byte
	i, j := len(a)-1, len(b)-1
	var over uint8 = 0
	for i >= 0 && j >= 0 {
		sum := (a[i] - '0') + (b[j] - '0') + over
		if sum >= 2 {
			over = 1
		} else {
			over = 0
		}
		ans = append(ans, sum%2+'0')
		i--
		j--
	}

	for ; i >= 0; i-- {
		sum := (a[i] - '0') + over
		if sum >= 2 {
			over = 1
		} else {
			over = 0
		}
		ans = append(ans, sum%2+'0')
	}

	for ; j >= 0; j-- {
		sum := (b[j] - '0') + over
		if sum >= 2 {
			over = 1
		} else {
			over = 0
		}
		ans = append(ans, sum%2+'0')
	}

	if over > 0 {
		ans = append(ans, '1')
	}

	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}

	return string(ans)
}
