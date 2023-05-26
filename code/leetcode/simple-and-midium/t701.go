package simple_and_midium

// 往二叉树中插入节点
//
//
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	pre := root
	tmpRoot := root
	for tmpRoot != nil {
		if val < tmpRoot.Val {
			pre = tmpRoot
			tmpRoot = tmpRoot.Left
		} else {
			pre = tmpRoot
			tmpRoot = tmpRoot.Right
		}
	}
	newNode := &TreeNode{Val: val}
	if val < pre.Val {
		pre.Left = newNode
	} else {
		pre.Right = newNode
	}
	return root
}
