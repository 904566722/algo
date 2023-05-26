package simple_and_midium

func calculate(s string) int {
	x, y := 1, 0
	A := func(x, y *int) {
		*x = 2 * (*x) + (*y)
	}
	B := func(x, y *int) {
		*y = 2 * (*y) + (*x)
	}
	for i := range s {
		if s[i] == 'A' {
			A(&x, &y)
		} else {
			B(&x, &y)
		}
	}
	return x + y
}