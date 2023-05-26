package simple_and_midium

func maximumXOR(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans |= num
	}
	return ans
}
