package simple_and_midium

func numberOfLines(widths []int, s string) []int {
	sum := 0
	cnt := 1
	for _, ch := range s {
		sum += widths[ch-'a']

		if sum > 100 {
			sum = widths[ch-'a']
			cnt++
		}
	}
	return []int{cnt, sum}
}
