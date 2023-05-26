package simple_and_midium

func findTheDifference(s string, t string) byte {
	st := s + t
	cnt := make([]int, 26)
	for i := range st {
		cnt[st[i]-'a']++
	}
	for i, v := range cnt {
		if v%2 == 1 {
			return 'a' + byte(i)
		}
	}
	return ' '
}
