package simple_and_midium

func rightSideView(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		ans = append(ans, q[len(q)-1].Val)
		n := len(q)
		for _, node := range q {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[n:]
	}
	return ans
}

// dfs 解法
// 用右节点优先的 rld
// 记录节点的深度
func rightSideView2(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	var st []*TreeNode
	tmpNode := root
	pre := root
	for len(st) > 0 || tmpNode != nil {
		for tmpNode != nil {
			st = append(st, tmpNode)
			if len(st) > len(ans) {
				ans = append(ans, tmpNode.Val)
			}
			tmpNode = tmpNode.Right
		}

		// 判断是否该访问该节点
		tmpNode = st[len(st)-1]
		if tmpNode.Left == nil || tmpNode.Left == pre {

			pre = tmpNode
			tmpNode = nil
			st = st[:len(st)-1]
		} else {
			tmpNode = tmpNode.Left
		}
	}
	return ans
}
