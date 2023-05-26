package simple_and_midium


func validPalindrome(s string) bool {
	l, r := 0, len(s) - 1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			flg1, flg2 := true, true
			// 去掉左边字符
			for i,j := l+1, r; i<j ; i,j = i+1, j-1 {
				if s[i] != s[j] {
					flg1 = false
					break
				}
			}
			// 去掉右边字符
			for i,j := l, r-1; i<j; i,j = i+1, j-1 {
				if s[i] != s[j] {
					flg2 = false
					break
				}
			}
			return flg1 || flg2
		}
	}
	return true
}