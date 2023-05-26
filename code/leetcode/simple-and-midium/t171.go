package simple_and_midium

func titleToNumber(columnTitle string) int {
	k := 1
	sum := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		sum += int(columnTitle[i]-'A'+1) * k
		k *= 26
	}
	return sum
}
