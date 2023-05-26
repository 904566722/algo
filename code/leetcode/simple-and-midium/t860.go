package simple_and_midium

func lemonadeChange(bills []int) bool {
	myMoney := make(map[int]int, len(bills))
	for _, bill := range bills {
		myMoney[bill]++
		if !enoughChange(myMoney, bill-5) {
			return false
		}
	}
	return true
}

// 优先使用大面值的钱找零
// 可能找零的面值有：15、5
// 15 的情况优先使用 10 + 5来找零
func enoughChange(myMoney map[int]int, bill int) bool {
	if bill == 0 {
		return true
	}
	switch bill {
	case 15:
		if myMoney[10] != 0 && myMoney[5] != 0 {
			myMoney[10]--
			myMoney[5]--
		} else if myMoney[5] >= 3 {
			myMoney[5] -= 3
		} else {
			return false
		}
	case 5:
		if myMoney[5] > 0 {
			myMoney[5]--
		} else {
			return false
		}
	}
	return true
}
