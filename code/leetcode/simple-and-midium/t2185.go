package simple_and_midium

import "strings"

func prefixCount(words []string, pref string) int {
	cnt := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			cnt ++
		}
	}
	return cnt
}