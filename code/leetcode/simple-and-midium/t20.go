package simple_and_midium

func isValid2(s string) bool {
	var stack []byte
	match := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for i := len(s) - 1; i >= 0; i-- {
		if len(stack) == 0 || match[s[i]] != stack[len(stack)-1] {
			stack = append(stack, s[i])
		} else if match[s[i]] == stack[len(stack)-1] {
			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack) == 0
}
