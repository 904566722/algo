package simple_and_midium

import "strings"

func removeOuterParentheses(s string) string {
	var stack []byte
	ans := ""
	tmp := strings.Builder{}
	for i := range s {
		if len(stack) == 0 || s[i] == '(' {
			stack = append(stack, s[i])
			tmp.WriteByte('(')
		} else if len(stack) > 0 && s[i] == ')' {
			stack = stack[0: len(stack)-1]
			tmp.WriteByte(')')
		}

		if len(stack) == 0 && tmp.Len() > 0 {
			tmps := tmp.String()
			ans += tmps[1:len(tmps)-1]
			tmp.Reset()
		}
	}
	return ans
}

func removeOuterParentheses2(s string) string {
	var stack, ans []byte
	for i := range s {
		ch := s[i]
		if ch == '(' {
			stack = append(stack, ch)
			if len(stack) > 1 {
				ans = append(ans, ch)
			}
		}
		if ch == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) != 0 {
				ans = append(ans, ch)
			}
		}
	}
	return string(ans)
}
