package simple_and_midium

import (
	"sort"
)

// 将 deck 分成许多子数组
// 		满足：子数组内元素均相同；每个子数组长度相同；长度大于等于2
func hasGroupsSizeX(deck []int) bool {
	sort.Ints(deck)
	deck = append(deck, -1)
	l := -1
	for i, j := 0, 1; j < len(deck) && deck[j] != -1; {
		for j < len(deck) && deck[j] == deck[i] {
			j++
		}
		if l == -1 {
			l = j - i
		} else if gcd(j-i, l) < 2 {
			return false
		}
		i = j
	}
	return l >= 2
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	} else {
		return gcd(b%a, a)
	}
}
