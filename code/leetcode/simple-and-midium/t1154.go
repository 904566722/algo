package simple_and_midium

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	daysMp := map[int]int {
		1: 31,
		3: 31,
		4: 30,
		5: 31,
		6: 30,
		7: 31,
		8: 31,
		9: 30,
		10: 31,
		11: 30,
		12: 31,
	}
	dateStrs := strings.Split(date, "-")
	daysMp[2] = 28
	y, _ := strconv.Atoi(dateStrs[0])
	m, _ := strconv.Atoi(dateStrs[1])
	d, _ := strconv.Atoi(dateStrs[2])
	if isLeapYear(y) {
		daysMp[2] = 29
	}
	cnt := 0
	for i := 1; i < m; i++ {
		cnt += daysMp[i]
	}
	return cnt + d
}

func isLeapYear(year int) bool {
	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
		return true
	}
	return false
}