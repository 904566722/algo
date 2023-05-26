package simple_and_midium

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	newMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		line := make([]int, m)
		newMatrix[i] = line
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			newMatrix[col][row] = matrix[row][col]
		}
	}
	return newMatrix
}
