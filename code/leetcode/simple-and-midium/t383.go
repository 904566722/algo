package simple_and_midium

func canConstruct(ransomNote string, magazine string) bool {
	mlen, rlen := len(magazine), len(ransomNote)
	if mlen < rlen {
		return false
	}

	mcnt := make([]int, 26)
	for i := range magazine {
		mcnt[magazine[i]-'a']++
	}
	for i := range ransomNote {
		mcnt[ransomNote[i]-'a']--
		if mcnt[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
