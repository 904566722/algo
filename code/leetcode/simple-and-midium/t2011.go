package simple_and_midium

func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, s := range operations {
		if s[1] == '+' {
			ans++
		} else {
			ans--
		}
	}
	return ans
}
