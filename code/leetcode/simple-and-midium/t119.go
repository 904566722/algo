package simple_and_midium

// 0 <= rowIndex <= 33
// so lenMax = 34

func getRow(rowIndex int) []int {
	var pre, cur []int
	// 计算每一行
	for i := 0; i <= rowIndex; i++ {
		// 每一行的计算方法
		cur = make([]int, i+1)
		cur[0] = 1
		cur[i] = 1
		for j := 1; j <= i-1; j++ {
			cur[j] = pre[j-1] + pre[j]
		}
		pre = cur
	}

	return cur
}
