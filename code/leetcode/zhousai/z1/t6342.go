package z1

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	lineCnt := make([]int, m)
	colCnt := make([]int, n)
	for i, num := range arr {
		LOOP:
		for r, l := range mat {
			for c, _ := range l {
				if num == mat[r][c] {
					lineCnt[r]++
					colCnt[c]++

					if lineCnt[r] == n {
						return i
					}
					if colCnt[c] == m {
						return i
					}
					break LOOP
				}
			}
		}
	}

	var ss []string
	for _, l := range mat {
		var s strings.Builder
		s.Grow(n)
		for _, num := range l {
			s.WriteString(strconv.Itoa(num))
		}
		ss = append(ss, s.String())
	}

	return -1
}

func firstCompleteIndex2(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	lineCnt := make([]int, m)
	colCnt := make([]int, n)

	num2idx := make([]int, m * n + 1)
	for i, l := range mat {
		for j, num := range l {
			num2idx[num] = n * i + j
		}
	}

	for i, num := range arr {
		idx := num2idx[num]

		row := idx/n
		col := idx%n
		lineCnt[row] ++
		colCnt[col]++

		if lineCnt[row] == n {
			return i
		}
		if colCnt[col] == m {
			return i
		}

	}

	return -1
}

func firstCompleteIndex3(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	lineCnt := make([]int, m)
	colCnt := make([]int, n)

	num2idx := make([]int, m * n + 2)

	preI := 0
	for i, num := range arr {
		idx := num2idx[num]
		// not exists
		if idx == 0 {
			LOOP:
			for _i := preI ; _i < m; _i++ {
				_j := 0
				//if _i == preI {
				//	_j = preJ
				//}
				for ; _j < n; _j++ {
					num2idx[mat[_i][_j]] = n * _i + _j + 1
					if mat[_i][_j] == num {
						preI = _i
						//preJ = _j
						idx = num2idx[num]
						break LOOP
					}
				}
			}
		}

		row := (idx-1)/n
		col := (idx-1)%n
		lineCnt[row] ++
		colCnt[col]++

		if lineCnt[row] == n {
			return i
		}
		if colCnt[col] == m {
			return i
		}

	}

	return -1
}

// 返回一个由 1~n 装填的一个二维数组
func makeData(m, n int) ([]int, [][]int) {
	// num range [1, m*n]
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, m * n + 1)
	for num := 1; num <= m * n; num++ {
		randPos := rand.Intn(m * n) // [0, n*m)
		for arr[randPos] != 0 {
			randPos = rand.Intn(m * n)
		}

		arr[randPos] = num
	}

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans[i][j] = arr[i*n + j]
		}
	}

	return arr, ans
}

func compEfficiency()  {
	befWinCnt, aftWinCnt := 0, 0
	for i := 1000; i <= 10000; i += 1000 {
		rand.Seed(time.Now().UnixNano())
		for j := 0; j < 10; j ++ {
			//nums, mat := makeData(i+j, i+j)
			_m, _n := i, i
			arr := make([]int, _m * _n + 1)
			rand.Seed(time.Now().UnixNano())
			for num := 1; num <= _m * _n; num++ {
				randPos := rand.Intn(_m * _n) // [0, n*m)
				for arr[randPos] != 0 {
					randPos = rand.Intn(_m * _n)
				}

				arr[randPos] = num
			}

			ans := make([][]int, _m)
			for i := range ans {
				ans[i] = make([]int, _n)
			}
			appeared := make([]int, _m * _n + 2)
			for i := 0; i < _m; i++ {
				for j := 0; j < _n; j++ {
					randnum := rand.Intn(_m * _n) + 1
					rand.Seed(time.Now().UnixNano())
					for appeared[randnum] == 1 {
						randnum = rand.Intn(_m * _n) + 1
					}
					ans[i][j] = randnum
					appeared[randnum] = 1
				}
			}




			nums := arr
			mat := ans
			fmt.Printf("---数据量：%d x %d \t 第%d次测试--- \n", i, i, j+1)
			start1 := time.Now()
			rst1 := firstCompleteIndex2(nums, mat)

			end1 := time.Now()
			rst2 := firstCompleteIndex3(nums, mat)

			end2 := time.Now()
			fmt.Printf("程序执行结果:\t优化前：%d\t优化后：%d\t\n程序执行时间:\t优化前:%s\t优化后:%s\t\n\n", rst1, rst2, end1.Sub(start1), end2.Sub(end1))
			if end1.Sub(start1) < end2.Sub(end1) {
				befWinCnt++
			} else if end1.Sub(start1) > end2.Sub(end1) {
				aftWinCnt++
			}
		}
	}
	fmt.Printf("优化前胜出：%d次\t优化后胜出：%d次\n", befWinCnt, aftWinCnt)
}

func testBreak() {
	i := 0
	loop:
	for ; i < 10; i++ {
		for j := 0; j < 10; j ++ {
			if j == 1 && i == 3 {
				break loop
			}
		}
	}
	fmt.Println(i)
}