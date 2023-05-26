package simple_and_midium

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	si := 0
	cnt := 0
	for i := 0; i < len(g); i++ {
		for si < len(s) {
			if g[i] <= s[si] {
				cnt++
				si++
				break
			}
			si++
		}
	}
	return cnt
}
