package simple_and_midium

func balancedStringSplit(s string) int {
	var st []byte
	cnt := 0
	for i := range s {
		if len(st) != 0 && st[len(st)-1] == s[i] || len(st) == 0 {
			st = append(st, s[i])
		} else {
			st = st[0:len(st)-1]
		}

		if len(st) == 0 {
			cnt++
		}
	}
	return cnt
}
