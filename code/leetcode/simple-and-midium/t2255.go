package simple_and_midium

import "strings"

func countPrefixes(words []string, s string) int {
	cnt := 0
	for _, w := range words {
		if strings.HasPrefix(s, w) {
			cnt++
		}
	}
	return cnt
}