package simple_and_midium

func isAlienSorted(words []string, order string) bool {
	// 首先要将 order 顺序映射成数值，来做比较
	orderNums := make(map[byte]int, len(order))
	for i := range order {
		orderNums[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		if compareStringByOrder(words[i-1], words[i], orderNums) == 1 {
			return false
		}
	}

	return true
}

func compareStringByOrder(s1, s2 string, order map[byte]int) int {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) && s1[i] == s2[j] {
		i++
		j++
	}
	if i == len(s1) && j == len(s2) {
		return 0
	} else if i == len(s1) {
		return 2
	} else if j == len(s2) {
		return 1
	}

	if order[s1[i]] > order[s2[j]] {
		return 1
	} else {
		return 2
	}
}
