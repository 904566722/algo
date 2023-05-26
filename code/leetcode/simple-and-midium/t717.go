package simple_and_midium

func isOneBitCharacter(bits []int) bool {
	lastLen := 0
	for i := 0; i < len(bits); {
		if bits[i] == 0 {
			i++
			lastLen = 1
		} else {
			i += 2
			lastLen = 2
		}
	}

	return lastLen == 1
}
