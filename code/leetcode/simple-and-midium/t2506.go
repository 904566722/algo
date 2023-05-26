package simple_and_midium

import (
	"sort"
	"strings"
)

func similarPairs(words []string) int {
	wordCnt := make(map[string]int)
	for _, w := range words {
		bs := []byte(w)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		wordCnt[rmDup(string(bs))]++
	}
	sum := 0
	for _, cnt := range wordCnt {
		if cnt > 0 {
			sum += (cnt * cnt - cnt) / 2
		}
	}
	return sum
}

func rmDup(s string) string {
	n := len(s)
	switch n {
	case 1:
		return s
	case 2:
		if s[0] == s[1] {
			return s[:1]
		} else {
			return s
		}
	}
	var ans strings.Builder
	ans.Grow(n)
	ans.WriteByte(s[0])
	preCh := s[0]
	for i := 1; i < n; i++ {
		if s[i] != preCh {
			ans.WriteByte(s[i])
			preCh = s[i]
		}
	}
	return ans.String()
}