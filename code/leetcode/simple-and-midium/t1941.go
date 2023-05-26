package simple_and_midium

func areOccurrencesEqual(s string) bool {
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}
	cntFreq := map[int]struct{}{}
	for _, v := range cnt {
		if v != 0 {
			cntFreq[v] = struct{}{}
		}
	}
	return len(cntFreq) == 1
}