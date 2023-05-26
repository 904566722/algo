package simple_and_midium

func hasPathSum(root *TreeNode, targetSum int) bool {
	return lrd(root, targetSum)
}

// lrd 后序遍历中，栈保存着从叶子节点到根节点路径上的节点
func lrd(root *TreeNode, tgtSum int) bool {
	if root == nil {
		return false
	}

	var st []*TreeNode
	tmpRoot := root
	pre := root
	for len(st) > 0 || tmpRoot != nil {

		// push
		for tmpRoot != nil {
			st = append(st, tmpRoot)
			tmpRoot = tmpRoot.Left
		}

		// get node
		tmpRoot = st[len(st)-1]
		// whether delete?
		if tmpRoot.Right == nil {
			// 说明是叶子节点，访问
			sum := 0
			for _, node := range st {
				sum += node.Val
			}
			if sum == tgtSum {
				return true
			}

			pre = tmpRoot
			tmpRoot = nil
			st = st[:len(st)-1]
		} else if tmpRoot.Right == pre {
			// 说明当前节点并不是叶子节点，但是右子树已经访问过
			pre = tmpRoot
			tmpRoot = nil
			st = st[:len(st)-1]
		} else {
			tmpRoot = tmpRoot.Right
		}
	}
	return false
}

func havePathToLeaf(root *TreeNode, tgtSum int) bool {
	// end
	if root.Left == nil && root.Right == nil {
		return root.Val == tgtSum
	}

	have1, have2 := false, false
	if root.Left != nil {
		have1 = havePathToLeaf(root.Left, tgtSum-root.Val)
	}
	if root.Right != nil {
		have2 = havePathToLeaf(root.Right, tgtSum-root.Val)
	}

	return have1 || have2
}
