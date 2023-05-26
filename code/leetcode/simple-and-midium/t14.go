package simple_and_midium

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	ans := strs[0]
	for true {
		i := 1
		for ; i < len(strs) && strings.HasPrefix(strs[i], ans); i++ {
		}
		if i >= len(strs) {
			return ans
		}
		ans = ans[:len(ans)-1]
		if ans == "" {
			return ""
		}
	}
	return ""
}
