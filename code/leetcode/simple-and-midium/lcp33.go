package simple_and_midium

import "math"

// 通过观察两种操作可以知道一个事实：
//	- 升级水桶的操作一定是在蓄水之前的
// A - 升级水桶； B - 蓄水
// AAA..BBBB.. 《--- 操作序列一定是这种情况
//
// Q1：蓄水操作需要多少次？
//	遍历两个数组，vat/bucket (如果不是整除，需要再 + 1) 	《--- 找到这个最大值
//
// 那么只需要遍历：升级水桶，求 Q1
// A 操作上限？最后必须有一个 B，也就是求 Q1 的过程中，如果最小值一直是 1，就不再遍历
func storeWater(bucket []int, vat []int) int {
	sum := 0
	for i := range vat {
		sum += vat[i]
	}
	if sum == 0 {
		return 0
	}
	ans := math.MaxInt
	for a := 0; ; a++ {
		b := math.MinInt
		flg := true // 用来标记是否 tmp 都为 1 或者都为 0
		fullJud := true
		for i := range bucket {
			tmp := 0
			// 如果不升级，水桶容量可能为 0
			if bucket[i] + a == 0 {
				flg = false
				fullJud = false
				break
			}
			if vat[i] % (bucket[i] + a) != 0 {
				tmp++
			}
			tmp += vat[i]/(bucket[i] + a)

			// 如果 B 的操作都为 1 或者 0，那么就可以退出循环了
			if tmp > 1 {
				flg = false
			}

			b = max(b, tmp)
		}

		if fullJud {
			ans = min(ans, b + a)
		}
		if flg {
			break
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b  {
		return b
	}
	return a
}