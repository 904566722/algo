package simple_and_midium

func powerfulIntegers(x int, y int, bound int) []int {
	var ans []int
	appeared := make([]int, bound + 1)
	if bound >= 2 {
		ans = append(ans, 2)
		appeared[2] = 1
	}

	if x == 1 && y == 1 {
		return ans
	}
	if x == 1 || y == 1 {
		notOne := x
		if x == 1 {
			notOne = y
		}

		z := 1
		for j := 0; ; j++ {
			sum := 1 + z * notOne
			z *= notOne
			if sum > bound {
				return ans
			}
			if appeared[sum] == 1 {
				continue
			}
			ans = append(ans, sum)
			appeared[sum] = 1
		}
	}

	xx, yy := 1, 1
	iMax := 0
	for i := 0; ; i++ {
		sum :=  xx * x + yy // x^ + y^0
		xx *= x
		if sum > bound {
			iMax = i + 1
			break
		}
		if appeared[sum] == 1 {
			continue
		}
		ans = append(ans, sum)
		appeared[sum] = 1
	}

	xx, yy = 1, 1
	for i := 0; i < iMax; i++ {
		for j := 1; ; j++ {
			sum := power(x, i) + power(y, j)
			if sum > bound {
				break
			}
			if appeared[sum] == 1 {
				continue
			}
			ans = append(ans, sum)
			appeared[sum] = 1
		}
	}
	return ans
}

func power(x, i int) int {
	if x == 1 {
		return 1
	}
	if i == 0 {
		return 1
	}
	xx := 1
	for i > 0 {
		xx *= x
		i--
	}
	return xx
}
