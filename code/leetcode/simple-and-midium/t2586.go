package simple_and_midium


func vowelStrings(words []string, left int, right int) int {
	vowels := map[byte]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
	}

	cnt := 0
	for i := left; i <= right; i++ {
		_, ok1 := vowels[words[i][0]]
		_, ok2 := vowels[words[i][len(words[i])-1]]
		if ok1 && ok2 {
			cnt++
		}
	}
	return cnt
}