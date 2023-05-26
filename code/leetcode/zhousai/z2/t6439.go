package z2

import "strings"

//
func minLength(s string) int {
	return cut(s)
}


func cut(s string) int {
	for {
		i := strings.Index(s, "AB")
		j := strings.Index(s, "CD")
		// 不能再切割
		if i < 0 && j < 0 {
			break
		}

		s = strings.ReplaceAll(s, "AB", "")
		s = strings.ReplaceAll(s, "CD", "")
	}

	return len(s)
}