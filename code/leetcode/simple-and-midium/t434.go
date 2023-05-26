package simple_and_midium

import "strings"

func countSegments(s string) int {
	strs := strings.Split(s, " ")
	cnt := 0
	for _, v := range strs {
		if len(v) != 0 {
			cnt++
		}
	}
	return cnt
}
