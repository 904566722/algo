package simple_and_midium

import "math"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var st []*TreeNode
	var path1, path2 []int
	tmpRoot := root
	pre := root
	for len(st) > 0 || tmpRoot != nil {
		for tmpRoot != nil {
			st = append(st, tmpRoot)
			tmpRoot = tmpRoot.Left
		}

		// get
		tmpRoot = st[len(st)-1]
		if tmpRoot.Right == nil || pre == tmpRoot.Right {
			// visit
			if tmpRoot.Val == p.Val {
				for _, node := range st {
					path1 = append(path1, node.Val)
				}
			}
			if tmpRoot.Val == q.Val {
				for _, node := range st {
					path2 = append(path2, node.Val)
				}
			}

			st = st[:len(st)-1]
			pre = tmpRoot
			tmpRoot = nil
		} else {
			tmpRoot = tmpRoot.Right
		}
	}

	valIdx := map[int]int{}
	idxMax := math.MinInt
	for i, val := range path1 {
		valIdx[val] = i
	}
	for _, val := range path2 {
		if i, ok := valIdx[val]; ok && idxMax < i {
			idxMax = i
		}
	}
	return &TreeNode{
		Val: path1[idxMax],
	}
}
