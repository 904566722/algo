package simple_and_midium


func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		return true
	}
	n := len(s)
	s1, s2 := s, s
	for i := 0; i < n / 2; i++ {
		s1 = s1[1:] + string(s1[0])
		s2 = string(s2[n-1]) + s2[0:n-1]

		if s1 == goal || s2 == goal {
			return true
		}
	}
	return false
}