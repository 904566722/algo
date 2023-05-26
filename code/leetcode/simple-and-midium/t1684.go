package simple_and_midium

func countConsistentStrings(allowed string, words []string) int {
	appeared := make([]int, 26)
	for i := range allowed {
		appeared[allowed[i]-'a'] = 1
	}

	cnt := 0
	for _, w := range words {
		flg := 1
		for i := range w {
			if appeared[w[i]-'a'] == 0 {
				flg = 0
				break
			}
		}
		cnt += flg
	}
	return cnt
}