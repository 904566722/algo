package simple_and_midium

import "strconv"

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[int]bool, 9)
	cols := make([]map[int]bool, 9)
	nines := make([]map[int]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = map[int]bool{}
		cols[i] = map[int]bool{}
		nines[i] = map[int]bool{}
	}

	for i, line := range board {
		for j := range line {
			if board[i][j] == '.' {
				continue
			}

			num, _ := strconv.Atoi(string(board[i][j]))
			if rows[i][num] {
				return false
			}
			rows[i][num] = true
			if cols[j][num] {
				return false
			}
			cols[j][num] = true
			if nines[ij2Idx(i, j)][num] {
				return false
			}
			nines[ij2Idx(i, j)][num] = true
		}
	}
	return true
}

func ij2Idx(i, j int) int {
	if i < 3 && j < 3 {
		return 0
	} else if i < 3 && j < 6 {
		return 1
	} else if i < 3 && j < 9 {
		return 2
	} else if i < 6 && j < 3 {
		return 3
	} else if i < 6 && j < 6 {
		return 4
	} else if i < 6 && j < 9 {
		return 5
	} else if j < 3 {
		return 6
	} else if j < 6 {
		return 7
	} else {
		return 8
	}
}