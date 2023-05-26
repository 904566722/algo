package simple_and_midium

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	var rst []string
	// 创建一个新的数组，将分数排序
	// 获得每种分数的排名 scoreRank k-v：分数-排名；
	scoreTmp := make([]int, len(score)) // deep copy
	copy(scoreTmp, score)

	scoreRank := make(map[int]int)
	sort.Sort(sort.Reverse(sort.IntSlice(scoreTmp))) // 降序排序
	for i, v := range scoreTmp {
		scoreRank[v] = i + 1
	}

	// 遍历分数，获得结果
	for _, v := range score {
		rst = append(rst, award(scoreRank[v]))
	}
	return rst
}

func award(n int) string {
	mp := map[int]string{
		1: "Gold Medal",
		2: "Silver Medal",
		3: "Bronze Medal",
	}
	if n <= 3 {
		return mp[n]
	} else {
		return strconv.Itoa(n)
	}
}
