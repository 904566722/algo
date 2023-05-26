package simple_and_midium

import (
	"math"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	otherSet := "!?',;."
	pbs := []byte(paragraph)
	for i := range paragraph {
		if strings.Contains(otherSet, string(paragraph[i])) {
			pbs[i] = ' '
		}
	}
	paragraph = string(pbs)
	words := strings.Split(paragraph, " ")
	cnt := make(map[string]int)
	for _, w := range words {
		cnt[strings.ToLower(w)]++
	}
	bannedStr := strings.Join(banned, " ")
	_max := math.MinInt
	ans := ""
	for w,num := range cnt {
		if strings.Contains(bannedStr, w) {
			continue
		}
		if _max < num {
			_max = num
			ans = w
		}
	}
	return ans
}
