package simple_and_midium

// 原问题求解：左子树与右子树 较大的高度+1
// 子问题：左子树与右子树 能够用相同的方法求解
// 结束条件：空节点，高度=0
func maxDepth(root *TreeNode) int {
	// end
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}
