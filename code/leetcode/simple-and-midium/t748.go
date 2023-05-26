package simple_and_midium

import (
	"math"
	"sort"
	"strings"
	"unicode"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlateTmp := rmDigitSpaceLower(licensePlate)
	minLen := math.MaxInt
	rst := ""
	for k := len(words) - 1; k >= 0; k-- {
		wordByte := []byte(strings.ToLower(words[k]))
		sort.Slice(wordByte, func(i, j int) bool {
			return wordByte[i] < wordByte[j]
		})

		i, j := 0, 0
		for i < len(licensePlateTmp) && j < len(wordByte) {
			if licensePlateTmp[i] == wordByte[j] {
				i++
			}
			j++
		}
		if i == len(licensePlateTmp) && minLen >= len(wordByte) {
			minLen = len(wordByte)
			rst = words[k]
		}
	}
	return rst
}

func rmDigitSpaceLower(s string) string {
	var rst []rune
	for _, b := range []rune(s) {
		if unicode.IsLetter(b) {
			rst = append(rst, unicode.ToLower(b))
		}
	}
	sort.Slice(rst, func(i, j int) bool {
		return rst[i] < rst[j]
	})
	return string(rst)
}
