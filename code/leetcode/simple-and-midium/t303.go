package simple_and_midium

//type NumArray struct {
//	Nums []int
//}
//
//func Constructor(nums []int) NumArray {
//	return NumArray{
//		nums,
//	}
//}
//
//func (this *NumArray) SumRange(left int, right int) int {
//	sum := 0
//	for i := left; i <= right; i++ {
//		sum += this.Nums[i]
//	}
//	return sum
//}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

type NumArray struct {
	Nums []int
	Sums []int
}

// Constructor 因为频繁调用到函数是 SumRange，而Constructor只调用一次，如果在该方法中预处理一下，就能减少SumRange的调用时间
//func Constructor(nums []int) NumArray {
//	// 将数组中存的数改为前i个数的和
//	// Sums[i] 表示 Nums[1] + Nums[2] + ... + Nums[i]
//	sums := nums
//	for i := 1; i < len(nums); i++ {
//		sums[i] = sums[i-1] + nums[i]
//	}
//
//	return NumArray{
//		Nums: nums,
//		Sums: sums,
//	}
//}
//
//func (this *NumArray) SumRange(left int, right int) int {
//	if left == 0 {
//		return this.Sums[right]
//	}
//	return this.Sums[right] - this.Sums[left-1]
//}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
