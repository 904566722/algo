package simple_and_midium

func fac(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	var n1, n2, f1, f2 int
	if n % 2 == 0 {
		n1, n2 = n/2, n/2
	} else {
		n1, n2 = (n-1)/2, (n+1)/2
	}

	f1 = fac(n1)
	f2 = fac(n2)
	return f1 * f2
}