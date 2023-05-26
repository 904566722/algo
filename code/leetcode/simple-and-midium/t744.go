package simple_and_midium

func nextGreatestLetter(letters []byte, target byte) byte {
	if int(target)-int(letters[len(letters)-1]) >= 0 {
		return letters[0]
	}
	for _, letter := range letters {
		if int(letter)-int(target) > 0 {
			return letter
		}
	}
	return letters[0]
}

// byte 直接比较
func nextGreatestLetter2(letters []byte, target byte) byte {
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	for _, letter := range letters {
		if letter > target {
			return letter
		}
	}
	return letters[0]
}
