package simple_and_midium

//func addStrings(num1 string, num2 string) string {
//	var ans []byte
//	n1, n2 := len(num1), len(num2)
//	minLen := string2.min(n1, n2)
//	i, sum := 0, byte(0)
//	for ; i < minLen; i++ {
//		sum += (num1[n1-i-1] - '0') + (num2[n2-i-1] - '0')
//		ans = append(ans, sum%10+'0')
//		sum /= 10
//	}
//	for ; i < n1; i++ {
//		sum += num1[n1-i-1] - '0'
//		ans = append(ans, sum%10+'0')
//		sum /= 10
//	}
//	for ; i < n2; i++ {
//		sum += num2[n2-i-1] - '0'
//		ans = append(ans, sum%10+'0')
//		sum /= 10
//	}
//	if sum > 0 {
//		ans = append(ans, sum+'0')
//	}
//	for i, n := 0, len(ans); i < n/2; i++ {
//		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
//	}
//	return string(ans)
//}
//
//// 代码优化
//func addStrings2(num1 string, num2 string) string {
//	var ans []byte
//	n1, n2 := len(num1), len(num2)
//	maxLen := string2.max(n1, n2)
//	i, sum := 0, byte(0)
//	for ; i < maxLen; i++ {
//		a, b := byte(0), byte(0)
//		if n1-i-1 >= 0 {
//			a = num1[n1-i-1] - '0'
//		}
//		if n2-i-1 >= 0 {
//			b = num2[n2-i-1] - '0'
//		}
//		sum += a + b
//		ans = append(ans, sum%10+'0')
//		sum /= 10
//	}
//	//for ; i < n1; i++ {
//	//	sum += num1[n1-i-1] - '0'
//	//	ans = append(ans, sum%10+'0')
//	//	sum /= 10
//	//}
//	//for ; i < n2; i++ {
//	//	sum += num2[n2-i-1] - '0'
//	//	ans = append(ans, sum%10+'0')
//	//	sum /= 10
//	//}
//	if sum > 0 {
//		ans = append(ans, sum+'0')
//	}
//	for i, n := 0, len(ans); i < n/2; i++ {
//		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
//	}
//	return string(ans)
//}
