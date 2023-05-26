package simple_and_midium

import "strings"

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	var ans []string
	for _, v := range strs {
		bs := []byte(v)
		for i, n := 0, len(bs); i<n/2; i++ {
			bs[i], bs[n-i-1] = bs[n-i-1], bs[i]
		}
		ans = append(ans, string(bs))
	}
	return strings.Join(ans, " ")
}