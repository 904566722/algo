package z1

func isWinner(player1 []int, player2 []int) int {
	s1, s2 := player1[0], player2[0]
	n := len(player1)
	switch n {
	case 1:
		if s1 > s2 {
			return 1
		} else if s1 < s2 {
			return 2
		} else {
			return 0
		}
	default:
		if s1 == 10 {
			s1 += 2 * player1[1]
		} else {
			s1 += player1[1]
		}
		if s2 == 10 {
			s2 += 2 * player2[1]
		} else {
			s2 += player2[1]
		}
	}

	for i := 2; i < n; i++ {
		if player1[i-1] == 10 || player1[i-2] == 10 {
			s1 += 2 * player1[i]
			//player1[i] *= 2
		} else {
			s1 += player1[i]
		}

		if player2[i-1] == 10 || player2[i-2] == 10 {
			s2 += 2 * player2[i]
			//player2[i] *= 2
		} else {
			s2 += player2[i]
		}
	}
	if s1 > s2 {
		return 1
	} else if s1 < s2 {
		return 2
	} else {
		return 0
	}
}
