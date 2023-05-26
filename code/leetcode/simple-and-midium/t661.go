package simple_and_midium

type neighborhood struct {
	row int
	col int
}

var nbs = []neighborhood{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 0}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func imageSmoother(img [][]int) [][]int {
	var rst [][]int
	for row, line := range img {
		rstLine := make([]int, len(line))
		rst = append(rst, rstLine)
		for col, _ := range line {
			sum := 0
			cnt := 0
			for _, nb := range nbs {
				r, c := row+nb.row, col+nb.col
				if r < 0 || r >= len(img) || c < 0 || c >= len(line) {
					continue
				}
				sum += img[r][c]
				cnt++
			}
			rst[row][col] = sum / cnt
		}
	}
	return rst
}
