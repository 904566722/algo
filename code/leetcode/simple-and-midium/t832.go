package simple_and_midium

func flipAndInvertImage(image [][]int) [][]int {
	var rst [][]int
	for _, line := range image {
		// 水平翻转，同时反转
		newLine := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			newLine[i] = 1 - line[len(line)-i-1]
		}
		rst = append(rst, newLine)
	}
	return rst
}
