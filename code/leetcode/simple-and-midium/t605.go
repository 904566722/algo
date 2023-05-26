package simple_and_midium

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); {
		if flowerbed[i] == 1 {
			i++
		} else if canPlanted(flowerbed, i-1, i+1) {
			flowerbed[i] = 1
			n--
		}
		if n <= 0 {
			break
		}
		i++
	}

	if n <= 0 {
		return true
	}
	return false
}

func canPlanted(flowerbed []int, left, right int) bool {
	leftFlg, rightFlg := false, false
	if left < 0 {
		leftFlg = true
	} else if flowerbed[left] == 0 {
		leftFlg = true
	}

	if right >= len(flowerbed) {
		rightFlg = true
	} else if flowerbed[right] == 0 {
		rightFlg = true
	}

	return leftFlg && rightFlg
}
