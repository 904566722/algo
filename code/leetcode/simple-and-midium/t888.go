package simple_and_midium

// 设 sum1、sum2 分别表示 alice、bob 的糖果总数
// a、b 表示 alice、bob 交换的糖果数量， sum1-a+b = sum2-b+a，2b-2a = sum2-sum1
// 问题转化为，找到两个数组中的两个数，其差值为总数差的一半
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	ans := make([]int, 2)
	sum1, sum2 := 0, 0
	aliceCnt := make(map[int]struct{})
	for i := 0; i < len(aliceSizes); i++ {
		sum1 += aliceSizes[i]
		aliceCnt[aliceSizes[i]*2] = struct{}{}
	}
	for _, v := range bobSizes {
		sum2 += v
	}

	for _, b := range bobSizes {
		twoA := b*2 - sum2 + sum1
		if _, ok := aliceCnt[twoA]; ok {
			ans[0] = twoA / 2
			ans[1] = b
			break
		}
	}

	return ans
}
