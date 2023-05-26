package simple_and_midium

func pathSum(root *TreeNode, targetSum int) [][]int {
	return lrdPath(root, targetSum)
}

func lrdPath(root *TreeNode, targetSum int) [][]int {
	var st []*TreeNode
	var ans [][]int
	if root == nil {
		return ans
	}

	tmpRoot := root
	pre := root
	for len(st) > 0 || tmpRoot != nil {
		for tmpRoot != nil {
			st = append(st, tmpRoot)
			tmpRoot = tmpRoot.Left
		}

		tmpRoot = st[len(st)-1]
		if tmpRoot.Right == nil || tmpRoot.Right == pre {
			// visit
			if tmpRoot.Left == nil && tmpRoot.Right == nil {
				sum := 0
				var path []int
				for _, node := range st {
					sum += node.Val
					path = append(path, node.Val)
				}
				if sum == targetSum {
					ans = append(ans, path)
				}
			}

			st = st[:len(st)-1]
			pre = tmpRoot
			tmpRoot = nil
		} else {
			tmpRoot = tmpRoot.Right
		}
	}
	return ans
}
