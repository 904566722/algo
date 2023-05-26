package simple_and_midium

//
//func backspaceCompare(s string, t string) bool {
//	n1, n2 := len(s), len(t)
//	n := string2.max(n1, n2)
//	var stack1, stack2 []byte
//	for i:=0; i<n; i++ {
//		idx1, idx2 := n1-i-1, n2-i-1
//		if idx1 >= 0 {
//			if len(stack1) != 0 && stack1[len(stack1)-1]=='#' && s[idx1] != '#' {
//				stack1 = stack1[0: len(stack1)-1]
//			} else {
//				stack1 = append(stack1, s[idx1])
//			}
//		}
//		if idx2 >= 0 {
//			if len(stack2) != 0 && stack2[len(stack2)-1]=='#' && t[idx2] != '#' {
//				stack2 = stack2[0: len(stack2)-1]
//			} else {
//				stack2 = append(stack2, t[idx2])
//			}
//		}
//	}
//	return strings.TrimRight(string(stack1), "#") == strings.TrimRight(string(stack2), "#")
//}
