package simple_and_midium

func minDeletionSize(strs []string) int {
	col := len(strs[0])
	cnt := 0
	for c := 0; c < col; c++ {
		r := 1
		for ; r < len(strs) && strs[r][c] > strs[r-1][c]; r++ {
		}
		if r != len(strs) {
			cnt++
		}
	}
	return cnt
}
