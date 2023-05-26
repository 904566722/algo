package simple_and_midium


func largeGroupPositions(s string) [][]int {
	var ans [][]int
	start, end, n := 0, 0, len(s)
	for start < n {
		for end < n && s[end] == s[start] { end++ }
		if end - start >= 3 {
			ans = append(ans, []int{start, end-1})
		}
		start = end
	}
	return ans
}