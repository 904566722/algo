package simple_and_midium

func greatestLetter(s string) string {
	freq := map[byte]int{}
	for _, ch := range s {
		freq[byte(ch)]++
	}
	for i := 'Z'; i >= 'A'; i-- {
		lower :=  i + 'a' - 'A'
		if freq[byte(i)] > 0 && freq[byte(lower)] > 0 {
			return string(i)
		}
	}
	return ""
}