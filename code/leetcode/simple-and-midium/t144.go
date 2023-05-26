package simple_and_midium

func preorderTraversal(root *TreeNode) []int {
	var nums []int
	pOrder(root, &nums)
	return nums
}

// DLR
func pOrder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	*nums = append(*nums, root.Val)
	pOrder(root.Left, nums)
	pOrder(root.Right, nums)

	return
}
