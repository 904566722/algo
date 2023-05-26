package simple_and_midium

import (
	"math"
)

//统计出现的字母的频率
//找到最大值、最小值
//
//期望：只有一个最大值，最小值与最大值差值为1 || 最大值为1
func equalFrequency(word string) bool {
	freq := map[byte]int{}

	for _, ch := range word {
		freq[byte(ch)]++
	}
	if len(freq) == 1 {
		return true
	}

	minV, maxV := math.MaxInt, math.MinInt
	freqCnt := map[int]int{}
	for _, v := range freq {
		if minV > v {
			minV = v
		}
		if maxV < v {
			maxV = v
		}
		freqCnt[v]++
	}

	if maxV - 1 == minV && freqCnt[maxV] == 1 {
		return true
	}
	if minV-1 == 0 && len(freqCnt) == 1 {
		return true
	}
	l := freqCnt[minV]
	if delete(freqCnt, minV); minV-1 == 0 && len(freqCnt) == 1 && l == 1{
		return true
	}

	return false
}