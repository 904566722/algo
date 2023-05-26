package simple_and_midium

func minNumberOfFrogs(croakOfFrogs string) int {
	cnt := make([]int, 5)
	for _, ch := range croakOfFrogs {
		switch ch {
		case 'c':
			if cnt[4] != 0 {
				cnt[4]--
			}
			cnt[0]++
		case 'r':
			if cnt[0] == 0 {
				return -1
			}
			cnt[0]--
			cnt[1]++
		case 'o':
			if cnt[1] == 0 {
				return -1
			}
			cnt[1]--
			cnt[2]++
		case 'a':
			if cnt[2] == 0 {
				return -1
			}
			cnt[2]--
			cnt[3]++
		case 'k':
			if cnt[3] == 0 {
				return -1
			}
			cnt[3]--
			cnt[4]++
		}
	}
	if cnt[0] + cnt[1] + cnt[2] + cnt[3] == 0 {
		return cnt[4]
	}
	return -1
}