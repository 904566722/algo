package simple_and_midium

// bfs
// 第一层 1    2的0次方
// 第二层 2    2的1次方
// ...
// 遍历每一层的的时候，计算下一层的节点数量，如果不等于 2 的 n 次方，判断当前层的子树情况
// 设定节点的几个状态：
// 有左右节点：    3
// 有左节点：      2
// 有右节点：      9
// 没有子节点：    0
// 那么：  3 后面必须是    3、2、0
//         2 后面必须是    0
//         0 后面必须是    0
//         如果包含 9 则不完全
// 或者说：    3 前面必须是    3
//             2 前面必须是    3
//             0 前面必须是    3 或者2
func isCompleteTree(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}

	var q []*TreeNode
	q = append(q, root)
	nextLevelNum := 2
	preState := 3 // 假设在前面补充一个状态最好的节点
	//curLevelFull := true
	for len(q) > 0 {
		width := len(q)
		childNum := 0
		for i := 0; i < width; i++ {
			node := q[i]
			nodeState := 3
			if node.Left != nil && node.Right != nil {
				nodeState = 3
				childNum += 2
				q = append(q, node.Left)
				q = append(q, node.Right)
			} else if node.Left != nil {
				nodeState = 2
				childNum += 1
				q = append(q, node.Left)
			} else if node.Right != nil {
				nodeState = 9
				childNum += 1
				q = append(q, node.Right)
			} else if node.Left == nil && node.Right == nil {
				nodeState = 0
			}

			//
			if nodeState == 9 {
				return false
			}
			if nodeState == 3 && preState != 3 {
				return false
			}
			if nodeState == 2 && preState != 3 {
				return false
			}
			if nodeState == 0 && preState == 3 || preState == 2 {
				//
			}

			preState = nodeState
		}

		//if childNum != nextLevelNum && !curLevelFull {
		//	return false
		//}
		//if childNum != nextLevelNum {
		//	curLevelFull = false
		//}
		nextLevelNum *= 2

		q = q[width:]
	}
	return true
}
