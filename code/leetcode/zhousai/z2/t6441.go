package z2

import (
	"strconv"
)

// 核心步骤：
//	子问题1. num, num*num 字符串是否满足：存在一种分割情况，使得 a+b+...+x = num
//
//
// notice: n<=1000， n*n <=1000 000
// num：从 1 到 100 0000 的数是否符合子问题1是固定的！
//
// 1 9 10 36
// 1 81 100 1296
//
//
// 谈论 36 的情况， 分割 1 ，子问题是：296 能否分割成36-1=35 --》 可以用递归来实现这个子问题
func punishmentNumber(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if expectN(i, strconv.Itoa(i*i)) {
			ans += i * i
		}
	}
	return ans
}

// 用来解决子问题1
// squ 是否存在分割情况，使得各部分的和 == num
func expectN(num int, squ string) bool {
	if squ == "" {
		return false
	}
	squN, _ := strconv.Atoi(squ)
	if num == squN {
		return true
	}

	for i := 1; i < len(squ); i++ {
		a, _ := strconv.Atoi(squ[:i])
		if expectN(num-a, squ[i:]) {
			return true
		}
	}
	return false
}