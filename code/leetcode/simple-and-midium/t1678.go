package simple_and_midium

import "strings"

func interpret(command string) string {
	command = strings.ReplaceAll(command, "(al)", "al")
	command = strings.ReplaceAll(command, "()", "o")
	return command
}

func interpret2(command string) string {
	ans := strings.Builder{}
	ans.Grow(len(command))
	for i, n := 0, len(command); i < n; {
		if command[i] == '(' {
			if command[i+1] == ')' {
				i += 2
				ans.WriteString("o")
			}  else {
				i += 4
				ans.WriteString("al")
			}
		} else {
			ans.WriteString(string(command[i]))
			i++
		}
	}
	return ans.String()
}