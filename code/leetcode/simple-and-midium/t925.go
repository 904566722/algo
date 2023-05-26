package simple_and_midium

// 要保证【顺序、次序】
func isLongPressedName(name string, typed string) bool {
	if name[0] != typed[0] || name[len(name)-1] != typed[len(typed)-1] ||
		len(name) > len(typed) {
		return false
	}
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if typed[j] == name[i] {
			i++
			j++
		} else if i-1 >= 0 && typed[j] == name[i-1] {
			j++
		} else {
			return false
		}
	}

	for ; j < len(typed); j++ {
		if typed[j] != name[len(name)-1] {
			return false
		}
	}
	return i >= len(name)
}