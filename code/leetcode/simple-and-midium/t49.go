package simple_and_midium

import "sort"

func groupAnagrams(strs []string) [][]string {
	group := map[string][]string{}
	for _, s := range strs {
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		group[string(bs)] = append(group[string(bs)], s)
	}
	var ans [][]string
	for _, v := range group {
		ans = append(ans, v)
	}
	return ans
}
