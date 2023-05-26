package simple_and_midium

import (
	"bytes"
)

func findWords(words []string) []string {
	var rst []string
	for _, word := range words {
		flg := true
		line := -1
		for _, b := range []byte(word) {
			tmpLine := belongLine(string(b))
			if line == -1 {
				line = tmpLine
			} else if tmpLine != line {
				flg = false
				break
			}
		}

		if flg {
			rst = append(rst, word)
		}
	}
	return rst
}

func belongLine(b string) int {
	firstLine := []byte("qwertyuiopQWERTYUIOP")
	secondLine := []byte("asdfghjklASDFGHJKL")
	thirdLine := []byte("zxcvbnmZXCVBNM")

	if bytes.ContainsAny(firstLine, b) {
		return 1
	} else if bytes.ContainsAny(secondLine, b) {
		return 2
	} else if bytes.ContainsAny(thirdLine, b) {
		return 3
	}
	return 0
}
