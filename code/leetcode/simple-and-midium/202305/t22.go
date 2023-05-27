package _02305

// 栈 + 回溯
// 使用单调栈来表示状态,使用回溯列举所有情况
// 当所有字符都放入栈中时,如果栈为空,则符合情况,否则不符合
// 剪枝:如果没有左括号了,则把右括号全部放入
//
// 假设 n = 2
// (())
// ((
func generateParenthesis(n int) []string {
	var path []byte
	ans := make([]string, 0)
	// 新增一个数组来保存当前状态左括号和右括号的数量
	xy := [2]uint8{0, 0}
	xy[0] = 1
	dlr(path, xy[:], '(', uint8(n), &ans)
	return ans
}

func dlr(path []byte, xy []uint8, ch byte, n uint8, ans *[]string) {
	path = append(path, ch)
	if len(path) == int(n*2) {
		*ans = append(*ans, string(path))
	}

	// 左括号的数量 < n 的时候，才能继续选择左括号作为下一个符号
	if xy[0] < n {
		xy[0]++
		dlr(path, xy, '(', n, ans)
		xy[0]--
	}
	// 只有当右括号 < 左括号，下一个才有可能选右括号
	if xy[1] < xy[0] {
		xy[1]++
		dlr(path, xy, ')', n, ans)
		xy[1]--
	}
	return
}

//func isValid(pat []byte) bool {
//	var stack []byte
//	for i := 0; i < len(pat); i++ {
//		if len(stack) != 0 && stack[len(stack)-1] == '(' && pat[i] == ')' {
//			stack = stack[:len(stack)-1]
//		} else {
//			stack = append(stack, pat[i])
//		}
//	}
//	return len(stack) == 0
//}
