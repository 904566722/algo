package simple_and_midium

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	ss1, ss2 := strings.Split(s1, " "), strings.Split(s2, " ")
	cnt :=  make(map[string]int)
	for _, w := range ss1 {
		cnt[w]++
	}
	for _, w := range ss2 {
		cnt[w]++
	}
	var ans []string
	for w, num := range cnt {
		if num == 1 {
			ans = append(ans, w)
		}
	}
	return ans
}

func uncommonFromSentences2(s1 string, s2 string) []string {
	cnt :=  make(map[string]int)

	// ** 两次一样的代码，做代码优化，包装成临时函数即可
	//for _, w := range ss1 {
	//	cnt[w]++
	//}
	//for _, w := range ss2 {
	//	cnt[w]++
	//}

	wordCnt := func(s string) {
		strs := strings.Split(s, " ")
		for _, w := range strs {
			cnt[w]++
		}
	}

	wordCnt(s1)
	wordCnt(s2)

	var ans []string
	for w, num := range cnt {
		if num == 1 {
			ans = append(ans, w)
		}
	}
	return ans
}