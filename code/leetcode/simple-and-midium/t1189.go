package simple_and_midium

import "math"

func maxNumberOfBalloons(text string) int {
	cnt := make([]int, 26)
	for i := range text {
		cnt[text[i] - 'a']++
	}
	cnt['l'-'a'] /= 2
	cnt['o'-'a'] /= 2
	m := math.MaxInt
	for _, ch := range "balon" {
		if m > cnt[ch - 'a'] {
			m = cnt[ch - 'a']
		}
	}
	return m
}