package simple_and_midium

// LDR
// 入栈直到最左
// 出栈，访问，右节点的最左节点
// 重复
func inorderTraversal(root *TreeNode) []int {
	tmpRoot := root
	var ans []int
	var st []*TreeNode
	for len(st) > 0 || tmpRoot != nil {

		for tmpRoot != nil {
			st = append(st, tmpRoot)
			tmpRoot = tmpRoot.Left
		}

		tmpRoot = st[len(st)-1]
		ans = append(ans, tmpRoot.Val)
		st = st[:len(st)-1]
		tmpRoot = tmpRoot.Right
	}
	return ans
}
