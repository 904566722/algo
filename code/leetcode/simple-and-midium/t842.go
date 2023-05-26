package simple_and_midium

import (
	"strings"
)

func toGoatLatin(sentence string) string {
	vowel := map[byte]bool {
		'a': true, 'e': true, 'i': true, 'o':true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O':true, 'U': true,
	}

	words := strings.Split(sentence, " ")
	ans := make([]string, len(words))
	for i, w := range words {
		newWord := w
		if _, ok := vowel[w[0]]; ok {
			newWord += "ma"
		} else {
			newWord = newWord[1:] + newWord[:1] + "ma"
		}
		for j := 0; j<i+1; j++ {
			newWord += "a"
		}
		ans[i] = newWord
	}
	return strings.Join(ans, " ")
}

var vowel = map[byte]bool {
'a': true, 'e': true, 'i': true, 'o':true, 'u': true,
'A': true, 'E': true, 'I': true, 'O':true, 'U': true,
}

func toGoatLatin2(sentence string) string {
	ans := strings.Builder{}
	n := len(sentence)
	no := 1
	for i := 0; i<n; {
		j := i
		ans.WriteString(" ")
		for j<n && sentence[j] != ' ' {j++}
		if _, ok := vowel[sentence[i]]; ok {
			ans.WriteString(sentence[i:j])
		} else {
			ans.WriteString(sentence[i+1:j])
			ans.WriteByte(sentence[i])
		}
		ans.WriteString("ma")
		ans.WriteString(strings.Repeat("a", no))
		no++
		i = j + 1
	}
	return ans.String()[1:]
}