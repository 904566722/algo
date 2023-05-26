package extra_study

func getNext(pat string) []int {
	n := len(pat)
	next := make([]int, n)

	for l, cur := 0, 1; cur < n; cur++ {
		// 不想等需要根据前面得到的 next，更新 l
		// 直到相等或者 l == 0
		for l > 0 && pat[l] != pat[cur] {
			l = next[l-1]
		}

		if pat[cur] == pat[l] {
			l++
		}
		next[cur] = l
	}

	return next
}

func KMP(s, pat string) int {
	n, m := len(s), len(pat)
	next := getNext(pat)
	i, j := 0, 0
	for i < n && j < m {
		for j > 0 && pat[j] != s[i] {
			j = next[j-1]
		}

		if s[i] == pat[j] {
			j++
		}
		i++
	}
	if j == m {
		return i - m
	} else {
		return -1
	}
}
