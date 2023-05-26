package simple_and_midium

import (
	"sort"
)

type Person struct {
	name   string // 16
	height int    // 8
}

func sortPeople(names []string, heights []int) []string {
	people := make([]Person, len(names))
	for i := range names {
		people[i] = Person{
			name:   names[i],
			height: heights[i],
		}
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].height > people[j].height
	})

	ans := make([]string, len(heights))
	for i, p := range people {
		ans[i] = p.name
	}
	return ans
}

//func codingP() {
//	fmt.Println(unsafe.Sizeof(Person{}))
//
//}

// 由于身高的范围是 [1,10^5]，并且约定了每个人的身高互不相同，
// 那么可以开辟一个 100001 的字符串数组，用来保存每个身高对应的姓名
// 最后倒序遍历这个字符串数组即可
// 时间复杂度 O(n) + O(100001) 空间复杂度：O(100001)
func sortPeople2(names []string, heights []int) []string {
	M := 100001
	heightNames := make([]string, M)
	for i, n := 0, len(heights); i < n; i++ {
		heightNames[heights[i]] = names[i]
	}
	var ans []string
	for i := len(heightNames) - 1; i >= 0; i-- {
		if heightNames[i] != "" {
			ans = append(ans, heightNames[i])
		}
	}
	return ans
}
