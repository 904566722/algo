package simple_and_midium

// 正常思路：平方后对数组排序，可以在原数组上操作，不需要额外空间，之后做排序，时间复杂度 o（nlogn）
//
// 有没有 o（n） 能实现的方法？
// 新开辟数组 ans 下标k，从后面开始遍历，
// 用 i，j 分别指向 nums 的两端，每次计算 nums[i] ^2, nums[j] ^2，ans[k] 取其中的大者（当i==j 的时候退出）
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j, k := 0, n-1, n-1
	for i != j {
		leftSquare := nums[i] * nums[i]
		rightSquare := nums[j] * nums[j]

		if leftSquare > rightSquare {
			ans[k] = leftSquare
			i++
		} else {
			ans[k] = rightSquare
			j--
		}
		k--
	}
	ans[k] = nums[i] * nums[i]
	return ans
}
