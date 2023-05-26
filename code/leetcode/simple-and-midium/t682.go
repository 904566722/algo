package simple_and_midium

import "strconv"

func calPoints(operations []string) int {
	var stack []int
	sum := 0
	for _, op := range operations {
		n := len(stack)
		switch op {
		case "+":
			score := stack[n-1] + stack[n-2]
			sum += score
			stack = append(stack, score)
		case "D":
			score := 2 * stack[n-1]
			sum += score
			stack = append(stack, score)
		case "C":
			sum -= stack[n-1]
			stack = stack[:n-1]
		default:
			score, _ := strconv.Atoi(op)
			sum += score
			stack = append(stack, score)
		}
	}
	return sum
}
