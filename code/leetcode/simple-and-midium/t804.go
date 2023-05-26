package simple_and_midium

func uniqueMorseRepresentations(words []string) int {
	morsTb := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.",
		"--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	mp := make(map[string]bool)
	for _, word := range words {
		mors := ""
		for _, ch := range word {
			mors += morsTb[ch-'a']
		}
		mp[mors] = true
	}

	return len(mp)
}
