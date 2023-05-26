package simple_and_midium

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	n := len(s)
	// 用来代表字符 i 从 s 到 t 到偏移量
	offset := make([]int8, 128)
	var dft int8 = -128
	scnt, tcnt, tcntFlg := 0, 0, make([]int, 128)
	for i := range offset {
		offset[i] = dft
	}
	for i := 0; i < n; i++ {
		ofs := int8(t[i] - s[i])
		// 字符偏移信息还没录入
		if offset[s[i]] == dft {
			scnt++
			offset[s[i]] = ofs
		} else if ofs != offset[s[i]] {
			return false
		}
		if tcntFlg[t[i]] == 0 {
			tcnt++
			tcntFlg[t[i]] = 1
		}
	}
	return tcnt == scnt
}
