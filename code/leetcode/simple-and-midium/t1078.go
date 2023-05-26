package simple_and_midium

import "strings"

func findOcurrences(text string, first string, second string) []string {
	ss := strings.Split(text, " ")
	var ans []string
	for j, n := 1, len(ss); j < n - 1; j++ {
		if ss[j] == second && ss[j-1] == first {
			ans = append(ans, ss[j+1])
		}
	}
	return ans
}
