package simple_and_midium

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	strs := strings.Split(strings.TrimRight(s, " "), " ")
	return len(strs[len(strs)-1])
}
