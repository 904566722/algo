package simple_and_midium

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	ruleKey2Idx := map[string]int{
		"type": 0,
		"color": 1,
		"name": 2,
	}

	ruleIdx := ruleKey2Idx[ruleKey]
	ans := 0
	for _, item := range items {
		if item[ruleIdx] == ruleValue {
			ans++
		}
	}
	return ans
}