package simple_and_midium

func reverseStr(s string, k int) string {
	n := len(s)
	bs := []byte(s)
	for l:=0; l < n; l += k {
		if l/k % 2 == 1 {
			continue
		}
		r := l + k
		if r > n {
			r = n
		}
		for cnt, i := 0, l; i < (l + r) / 2; i++ {
			bs[i], bs[r-cnt-1] = bs[r-cnt-1], bs[i]

			cnt++
		}
	}
	return string(bs)
}