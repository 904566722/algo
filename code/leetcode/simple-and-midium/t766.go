package simple_and_midium

func isToeplitzMatrix(matrix [][]int) bool {
	// 遍历第一行 r = 0
	for c := 0; c < len(matrix[0]); c++ {
		i := 0 + 1
		j := c + 1
		for i < len(matrix) && j < len(matrix[0]) && matrix[i][j] == matrix[0][c] {
			i, j = i+1, j+1
		}

		if i != len(matrix) && j != len(matrix[0]) {
			return false
		}
	}
	// 遍历第一列 c = 0
	for r := 0; r < len(matrix); r++ {
		i := r + 1
		j := 0 + 1
		for i < len(matrix) && j < len(matrix[0]) && matrix[i][j] == matrix[r][0] {
			i, j = i+1, j+1
		}

		if i != len(matrix) && j != len(matrix[0]) {
			return false
		}
	}

	return true
}
