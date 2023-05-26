package simple_and_midium

import "strings"

func isValid(s string) bool {
	idx := strings.Index(s, "abc")
	for idx >= 0 {
		if idx == 0 {
			s = s[3:]
		} else {
			s = s[0:idx] + s[idx+3:]
		}

		idx = strings.Index(s, "abc")
	}
	return len(s) == 0
}
