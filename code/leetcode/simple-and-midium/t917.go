package simple_and_midium

import (
	"unicode"
)

func reverseOnlyLetters(s string) string {
	i, j := 0,len(s)-1
	bs := []byte(s)
	for i < j {
		for i < len(s) && !unicode.IsLetter(rune(bs[i])) {i++}
		for j >= 0 && !unicode.IsLetter(rune(bs[j])) {j--}
		if i >= j {
			break
		}

		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
