package simple_and_midium

func spiralOrder(matrix [][]int) []int {
	// 移动方式
	moveId := 0
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, n * m)
	idx := 0
	I, J := moveId2ij(moveId)
	leftBoundary, upBoundary, rightBoundary, downBoundary := -1, -1, n, m
	i, j := 0, 0
	for idx < m * n {
		I, J = moveId2ij(moveId)
		for j < rightBoundary {
			ans[idx] = matrix[i][j]
			idx++
			i, j = i + I, j + J
		}
		if idx >= m * n {
			break
		}
		j--
		i++
		upBoundary++

		moveId++
		I, J = moveId2ij(moveId)
		for i < downBoundary {
			ans[idx] = matrix[i][j]
			idx++
			i, j = i + I, j + J
		}
		if idx >= m * n {
			break
		}
		i--
		j--
		rightBoundary--

		moveId++
		I, J = moveId2ij(moveId)
		for j > leftBoundary {
			ans[idx] = matrix[i][j]
			idx++
			i, j = i + I, j + J
		}
		if idx >= m * n {
			break
		}
		j++
		i--
		downBoundary--

		moveId++
		I, J = moveId2ij(moveId)
		for i > upBoundary {
			ans[idx] = matrix[i][j]
			idx++
			i, j = i + I, j + J
		}
		if idx >= m * n {
			break
		}
		i++
		j++
		leftBoundary++

		moveId = 0
	}
	return ans
}

func moveId2ij(moveId int) (int, int) {
	switch moveId{
	case 0:
		return 0, 1
	case 1:
		return 1, 0
	case 2:
		return 0, -1
	case 3:
		return -1, 0
	}
	return 0, 0
}
