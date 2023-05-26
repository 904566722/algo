package simple_and_midium

// 超时
func removeDuplicates(s string) string {
	if len(s) <= 1 {
		return s
	}

	for i,j := 0, 1; j < len(s); {
		if s[j] != s[i] {
			i++
			j++
		} else {
			return removeDuplicates(s[:i]+s[j+1:])
		}
		if j == len(s) {
			return s
		}
	}
	return ""
}

func removeDuplicates2(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	for i, j := 0, 1; j<n; n = len(s) {
		if s[i] == s[j] {
			s = s[:i] + s[j+1:]
			if i--; i<0 {
				i = 0
			}
			if j--; j<1 {
				j = 1
			}
		} else {
			i++
			j++
		}
	}
	return s
}

func removeDuplicates3(s string) string {
	var st []byte
	for i := range s {
		if len(st) != 0 && st[len(st)-1] == s[i] {
			st = st[:len(st)-1]
		} else {
			st = append(st, s[i])
		}
	}
	return string(st)
}

