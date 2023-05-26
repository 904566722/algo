package simple_and_midium

import "bytes"

func reverseVowels(s string) string {
	i, j, n := 0, len(s)-1, len(s)
	bs := []byte(s)

	for true {
		for ; i < n && !isVowel(bs[i]); i++ {
		}
		for ; j >= 0 && !isVowel(bs[j]); j-- {
		}
		if i >= j {
			break
		}
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}

	return string(bs)
}

func isVowel(b byte) bool {
	set := []byte{
		'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U',
	}
	return bytes.ContainsAny(set, string(b))
}
