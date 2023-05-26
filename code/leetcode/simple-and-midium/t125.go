package simple_and_midium

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for ; l < len(s) && !matchType(s[l]); l++ {
		}
		for ; r >= 0 && !matchType(s[r]); r-- {
		}

		if l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}

func matchType(b byte) bool {
	return unicode.IsLetter(rune(b)) || unicode.IsDigit(rune(b))
}
