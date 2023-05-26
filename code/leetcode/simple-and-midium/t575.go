package simple_and_midium

import (
	"math"
	"sort"
)

func distributeCandies(candyType []int) int {
	sort.Ints(candyType)
	cnt := 0
	lastType := math.MinInt
	for i := 0; i < len(candyType); i++ {
		if lastType != candyType[i] {
			lastType = candyType[i]
			cnt++
		}
	}
	if cnt >= len(candyType)/2 {
		return len(candyType) / 2
	} else {
		return cnt
	}
}

func distributeCandies2(candyType []int) int {
	mp := make(map[int]struct{})
	cnt := 0
	for _, t := range candyType {
		if _, ok := mp[t]; !ok {
			cnt++
			mp[t] = struct{}{}
		}
	}

	if cnt >= len(candyType)/2 {
		return len(candyType) / 2
	} else {
		return cnt
	}
}
