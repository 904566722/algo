package simple_and_midium

import "strings"

func generateTheString(n int) string {
	if n % 2 == 1 {
		return strings.Repeat("a", n)
	}
	return strings.Repeat("a", n - 1) + "b"
}