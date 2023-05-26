package simple_and_midium

// 编号 [1,n] 映射为下标  0..n-1
// a[] 数组下标代表信任的出发点的编号；b 数组下标代表被信任的人的标号，数值代表被新人的次数
// 找到 a 数组中数值为 0 的下标的编号集合，再在这个集合里面找到一个下标，这个下标在 b 中的数值为 n - 1(被除开自己以外的所有人信任)
func findJudge(n int, trust [][]int) int {
	a := make([]int, n)
	b := make([]int, n)
	for _, t := range trust {
		a[t[0]-1] = 1
		b[t[1]-1]++
	}

	for i := range a {
		if a[i] == 0 && b[i] == n-1 {
			return i + 1
		}
	}

	return -1
}

// 简化为一个数组，flg[i]++：i 被人相信，flg[i]--:i 相信别人
// 法官的 flg == n - 1
func findJudge2(n int, trust [][]int) int {
	flg := make([]int, n+1)
	for _, t := range trust {
		flg[t[0]]--
		flg[t[1]]++
	}

	for i := range flg {
		if flg[i] == n-1 && i != 0 {
			return i
		}
	}

	return -1
}
