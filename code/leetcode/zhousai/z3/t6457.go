package z3

func removeTrailingZeros(num string) string {
	n := len(num)
	if num[0] == '0' {
		return ""
	}
	end := n - 1
	for num[end] == '0' {
		end--
	}
	return num[:end+1]
}
