package simple_and_midium

import "sort"

// p      xxxx
// s xxxxxxxxxxxxxx
//        i  j
// 一个固定长度的窗口一直往右滑动就行了
// 关键在于对于异位词的判断
// 方法1. 使用字符计数的方法，需要判断26个字母的计数是否相同
// 方法2.
//      使用一个临时 tmpP 来表示
//          窗口滑动的时候，新进来的字符如果在 tmpP 中存在，就将 tmpP 中对应的字符删除，
//          并且把 flg[i] 标记为 1；否则标记为 0
//      那么如果是异位词的时候，tmpP == ""
//      窗口滑动的时候，如果离开窗口的字符的标记 flg[i] == 1，那么需要将这个字符重新加到 tmpP 中
//
func findAnagrams(s string, p string) []int {
	var ans []int
	m, n := len(p), len(s)
	flg := make([]int, n)
	tmp := &tmpP{
		chs: map[byte]int{},
		len: 0,
	}
	// 初始化 tmpP
	for i := range p {
		tmp.chs[p[i]]++
		tmp.len++
	}
	for i := 0; i < m; i++ {
		if tmp.chs[s[i]] > 0 {
			tmp.cut(s[i])
			flg[i] = 1
		}
	}

	if tmp.len == 0 {
		ans = append(ans, 0)
	}
	for j := m; j < n; j ++ {
		passIdx := j - m
		if flg[passIdx] == 1 {
			tmp.add(s[passIdx])
		}
		if tmp.chs[s[j]] > 0 {
			tmp.cut(s[j])
			flg[j] = 1
		}
		if tmp.len == 0 {
			ans = append(ans, passIdx + 1)
		}
	}
	return ans
}

type tmpP struct {
	chs map[byte]int
	len int
}

func (t *tmpP) cut(ch byte) {
	t.chs[ch]--
	t.len--
}

func (t *tmpP) add(ch byte)  {
	t.chs[ch]++
	t.len++
}

func findAnagrams2(s string, p string) []int {
	var ans []int
	m, n := len(p), len(s)
	pSort := []byte(p)
	sort.Slice(pSort, func(i, j int) bool {
		return pSort[i] < pSort[j]
	})
	
	for i := m - 1; i < n; i++ {
		tmp := []byte(s[i-m+1:i+1])
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		if string(tmp) == string(pSort) {
			ans = append(ans, i-m+1)
		}
	}
	return ans
}