package simple_and_midium

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	ci := 0
	var rst [][]int
	newLine := make([]int, c)
	for _, line := range mat {
		for _, v := range line {
			newLine[ci] = v
			ci++
			if ci >= c {
				ci = 0
				rst = append(rst, newLine)
				newLine = make([]int, c)
			}
		}
	}
	return rst
}

func matrixReshape2(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}

	var rst [][]int
	for i := 0; i < r; i++ {
		line := make([]int, c)
		rst = append(rst, line)
	}

	for x := 0; x < r*c; x++ {
		rst[x/c][x%c] = mat[x/n][x%n]
	}

	return rst
}
