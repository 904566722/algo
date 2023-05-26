package simple_and_midium


//首字母大写
//	其他字母均小写 或者均 大写 是正确的
//首字母小写
//	其他字母均小写 是正确的
func detectCapitalUse(word string) bool {
	if 'A' <= word[0] && word[0] <= 'Z' && pureType(word[1:]) >= 0 {
		return true
	}
	if 'a' <= word[0] && word[0] <= 'z' && (pureType(word[1:]) == 0 || len(word[1:]) == 0) {
		return true
	}
	return false
}

// 2 空字符
// -1 不是纯大写或纯小写
// 0 纯小写
// 1 纯大写
func pureType(s string) int {
	if len(s) == 0 {
		return 2
	}
	lower, upper := 0, 0
	for i := range s {
		if 'a' <= s[i] && s[i] <= 'z' {
			lower = 1
		}
		if 'A' <= s[i] && s[i] <= 'Z' {
			upper = 1
		}
		if lower + upper > 1 {
			return -1
		}
	}
	if lower == 1 {
		return 0
	}
	return 1
}