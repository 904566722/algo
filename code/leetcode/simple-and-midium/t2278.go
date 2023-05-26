package simple_and_midium

func percentageLetter(s string, letter byte) int {
	cnt := 0
	for i := range s {
		if s[i] == letter {
			cnt ++
		}
	}
	return 100 * cnt / len(s)
}