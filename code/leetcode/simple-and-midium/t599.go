package simple_and_midium

import "math"

// 目标：名字相同 & 索引和最小
func findRestaurant(list1 []string, list2 []string) []string {
	nameMp := make(map[string]int) // k:v -> name:index
	for i, name := range list1 {
		nameMp[name] = i
	}

	minIdxSum := math.MaxInt
	idxSumMp := make(map[string]int)

	// 下面的代码能够优化
	var tgtName []string
	for i, name := range list2 {
		if _, ok := nameMp[name]; ok {
			idxSumMp[name] = nameMp[name] + i
			if minIdxSum > idxSumMp[name] {
				minIdxSum = idxSumMp[name]
			}
		}
	}
	for k, v := range idxSumMp {
		if v == minIdxSum {
			tgtName = append(tgtName, k)
		}
	}

	return tgtName
}
