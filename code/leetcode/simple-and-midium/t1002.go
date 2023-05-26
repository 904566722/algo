package simple_and_midium

import (
	"fmt"
	"math"
	"reflect"
)

// ** 理解出错了
// 遍历每个字符串，做字符的去重操作，并对出现的字符计数 cnt
// target：cnt[i] == len(words)
func commonChars(words []string) []string {
	cnt := make(map[byte]int, 26)
	n := len(words)
	var ans []string
	for _, w := range words {
		w = rmDuplication(w)
		for i := range w {
			cnt[w[i]]++
			if cnt[w[i]] == n {
				ans = append(ans, string(w[i]))
			}
		}
	}
	return ans
}

func rmDuplication(s string) string {
	set := make(map[byte]bool)
	for i := range s {
		set[s[i]] = true
	}
	var ans []byte
	for k := range set {
		ans = append(ans, k)
	}
	return string(ans)
}

// flg[][] 为 n * 26 的二维数组，flg[i] 保存了 words[i] 中每个字符出现的次数
// 遍历 flg 每一列，找到最小值 m，如果 m == 0，说明有字符串没有这个字符，否则，往答案中加入 m 个该字符
//
// 优化：那么其实可以直接遍历 26 个字符，查找这个字符在每个字符中出现的次数的最小值，如果大于零则加入结果集
// 总结，该题目是要寻找一个字符集，所有的字母也才26个，直接从这个方向入手，别考虑太多
func commonChars2(words []string) []string {
	n := len(words)
	flg := make([][26]int, n)

	for r, w := range words {
		for j := range w {
			flg[r][w[j]-'a']++
		}
	}

	var ans []string
	for col := 0; col < 26; col++ {
		m := math.MaxInt
		for r := 0; r < n; r++ {
			m = min(m, flg[r][col])
		}
		for ; m > 0; m-- {
			ans = append(ans, string('a'+byte(col)))
		}
	}
	return ans
}

func codingP() {
	// byte -> uint8
	// byte(col) + 'a'
	i := byte(0)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)
}
