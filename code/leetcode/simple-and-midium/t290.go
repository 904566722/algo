package simple_and_midium

import "strings"

// 使用两个哈希表，来代表两边的对应关系
func wordPattern(pattern string, s string) bool {
	p2s := make(map[byte]string)
	s2p := make(map[string]byte)
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}
	for i := range pattern {
		if (p2s[pattern[i]] != "" && p2s[pattern[i]] != strs[i]) ||
			(s2p[strs[i]] != 0 && s2p[strs[i]] != pattern[i]) {
			return false
		}
		p2s[pattern[i]] = strs[i]
		s2p[strs[i]] = pattern[i]
	}
	return true
}
