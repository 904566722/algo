package simple_and_midium

// 使用递归实现
// 原问题的解决方法：
//		找到已经序列化的左子树的最右下的节点，让其连接根节点的右节点（右树也需要序列化）
//		让根节点的右节点为左子树的起始节点
// 子问题：
//		左右子树的序列化（解决方法同上）
//
//		原问题与子问题的关系，原问题需要用到子问题的结果，因此
//		上述的处理步骤需要在递归之后做
//	结束条件：
//		根节点为空，已经序列化，直接返回
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	leftTree := root.Left
	leftTreeLast := root.Left
	for leftTreeLast != nil && leftTreeLast.Right != nil {
		leftTreeLast = leftTreeLast.Right
	}
	if leftTree != nil {
		leftTreeLast.Right = root.Right
		root.Right = leftTree
	}
	root.Left = nil
}
