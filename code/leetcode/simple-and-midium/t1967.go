package simple_and_midium

import "strings"

func numOfStrings(patterns []string, word string) int {
	cnt := 0
	for _, p := range patterns {
		if strings.Contains(word, p) {
			cnt++
		}
	}
	return cnt
}