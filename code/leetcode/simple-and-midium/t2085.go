package simple_and_midium


func countWords(words1 []string, words2 []string) int {
	getFreq := func(words []string) map[string]int {
		freq := map[string]int{}
		for _, word := range words{
			freq[word]++
		}
		return freq
	}
	freq1 := getFreq(words1)
	freq2 := getFreq(words2)

	cnt := 0
	for k, v := range freq1 {
		if v == 1 && freq2[k] == 1 {
			cnt++
		}
	}
	return cnt
}