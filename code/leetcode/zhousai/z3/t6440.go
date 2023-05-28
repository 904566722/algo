package z3

func differenceOfDistinctValues(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])
	// dp1、dp2 分别表示左上角、右下角对角线上的元素集合
	dp1 := make([][]map[uint8]struct{}, m)
	dp2 := make([][]map[uint8]struct{}, m)
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		dp1[i] = make([]map[uint8]struct{}, n)
		dp2[i] = make([]map[uint8]struct{}, n)
		ans[i] = make([]int, n)
	}

	for r1, r2 := 0, m-1; r1 < m; {
		for c1, c2 := 0, n-1; c1 < n; {
			dp1[r1][c1] = map[uint8]struct{}{}
			dp2[r2][c2] = map[uint8]struct{}{}
			// 左上角元素是否存在
			if r1-1 >= 0 && c1-1 >= 0 {
				// 当前元素加入集合
				deepCopyMap(dp1[r1][c1], dp1[r1-1][c1-1])
			}
			dp1[r1][c1][uint8(grid[r1][c1])] = struct{}{}

			// 右下角元素是否存在
			if r2+1 < m && c2+1 < n {
				deepCopyMap(dp2[r2][c2], dp2[r2+1][c2+1])
			}
			dp2[r2][c2][uint8(grid[r2][c2])] = struct{}{}

			c1++
			c2--
		}
		r1++
		r2--
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			x, y := 0, 0
			if (r-1 >= 0 && c-1 >= 0) && (r+1 < m && c+1 < n) {
				x, y = len(dp1[r-1][c-1]), len(dp2[r+1][c+1])
			} else if r+1 < m && c+1 < n {
				// 上面一排，左边一列
				y = len(dp2[r+1][c+1])
			} else if r-1 >= 0 && c-1 >= 0 {
				x = len(dp1[r-1][c-1])
			}
			ans[r][c] = abs(x, y)
		}
	}
	return ans
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func deepCopyMap(m1, m2 map[uint8]struct{}) {
	for k := range m2 {
		m1[k] = struct{}{}
	}
}

func differenceOfDistinctValues2(grid [][]int) [][]int {
	// 根据对角线来遍历
	// 设从 (0,0) 开始的对角线左下方为 区域 1，其他为区域 2

	m := len(grid)
	n := len(grid[0])
	dp1 := make([][]uint8, m)
	dp2 := make([][]uint8, m)
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		dp1[i] = make([]uint8, n)
		dp2[i] = make([]uint8, n)
		ans[i] = make([]int, n)
	}

	// 先遍历区域 1 的情况
	// l 表示一条对角线的 r 起点
	for l := m - 1; l > 0; l-- {
		set1 := make(map[uint8]struct{})
		r, c := l, 0
		for ; r < m && c < n; r, c = r+1, c+1 {
			dp1[r][c] = uint8(len(set1))
			set1[uint8(grid[r][c])] = struct{}{}
		}
		// 倒序遍历，统计右下角的情况
		set1 = nil
		set2 := make(map[uint8]struct{})
		for r2, c2 := r-1, c-1; r2 >= l && c2 >= 0; r2, c2 = r2-1, c2-1 {
			dp2[r2][c2] = uint8(len(set2))
			set2[uint8(grid[r2][c2])] = struct{}{}
		}
	}
	// 区域2
	// l 表示一条对角线 c 的起点
	for l := 0; l < n; l++ {
		set1 := make(map[uint8]struct{})
		r, c := 0, l
		for ; r < m && c < n; r, c = r+1, c+1 {
			dp1[r][c] = uint8(len(set1))
			set1[uint8(grid[r][c])] = struct{}{}
		}
		// 倒序遍历，统计右下角的情况
		set1 = nil
		set2 := make(map[uint8]struct{})
		for r2, c2 := r-1, c-1; r2 >= 0 && c2 >= l; r2, c2 = r2-1, c2-1 {
			dp2[r2][c2] = uint8(len(set2))
			set2[uint8(grid[r2][c2])] = struct{}{}
		}
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			ans[r][c] = abs(int(dp1[r][c]), int(dp2[r][c]))
		}
	}
	return ans
}
