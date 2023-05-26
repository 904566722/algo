package z2

// 用结果来推过程
// 从中间往两边开始对称遍历，
//		如果左边的字符比较小，将右边的字符换成左边的字符
//		如果右边比较小，将左边的换成右边的字符
//		相等不做操作
//	n 奇数:	5 x x x x x
//				  n/2 = 2
//				l = n/2-1, r = n/2+1
//	n 偶数：6  x x x x x x
//				  2 3
//				 l = n/2-1, r = n/2
func makeSmallestPalindrome(s string) string {
	n := len(s)
	var l, r int
	if n%2 == 1 {
		r = n/2 + 1
	} else {
		r = n / 2
	}
	l = n/2 - 1

	bs := []byte(s)
	for l >= 0 && r < n {
		if bs[l] < bs[r] {
			bs[r] = bs[l]
		} else if bs[l] > bs[r] {
			bs[l] = bs[r]
		}

		l--
		r++
	}
	return string(bs)
}
