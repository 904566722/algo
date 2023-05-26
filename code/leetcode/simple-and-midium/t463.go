package simple_and_midium

func islandPerimeter(grid [][]int) int {
	rst := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 1 {
				rst += 4
				if row-1 >= 0 && grid[row-1][col] == 1 {
					rst -= 1
				}
				if col-1 >= 0 && grid[row][col-1] == 1 {
					rst -= 1
				}
				if col+1 < len(grid[row]) && grid[row][col+1] == 1 {
					rst -= 1
				}
				if row+1 < len(grid) && grid[row+1][col] == 1 {
					rst -= 1
				}
			}
		}
	}
	return rst
}
