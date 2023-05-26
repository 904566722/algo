package simple_and_midium

import (
	"math"
	"unicode"
)

func reformat(s string) string {
	n := len(s)
	bs := make([]byte, n)
	l1, l2 := 0, 0
	var nums []byte
	var letters []byte
	for i, ch := range s {
		if unicode.IsLetter(ch) {
			l2++
			letters = append(letters, s[i])
		} else {
			l1++
			nums = append(nums, s[i])
		}

		if l1 > n/2 + 1 || l2 > n/2 + 1 {
			return ""
		}
	}
	if math.Abs(float64(l1 - l2)) > 1 {
		return ""
	}
	idx1, idx2 := 0, 0
	if l1 > l2 {
		idx2 = 1
	} else {
		idx1 = 1
	}
	for i := 0 ; idx1 < n && i < l1; i++ {
		bs[idx1] = nums[i]
		idx1 += 2
	}
	for i := 0; idx2 < n && i < l2; i++ {
		bs[idx2] = letters[i]
		idx2 += 2
	}
	return string(bs)
}