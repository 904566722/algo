package simple_and_midium

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var q []*TreeNode
	q = append(q, root)

	i := 0
	for len(q) > 0 {
		var tmpQ []*TreeNode
		var vals []int
		for len(q) > 0 {
			tmp := q[0]
			q = q[1:]
			vals = append(vals, tmp.Val)

			if tmp.Left != nil {
				tmpQ = append(tmpQ, tmp.Left)
			}
			if tmp.Right != nil {
				tmpQ = append(tmpQ, tmp.Right)
			}
		}

		if i%2 == 1 {
			for idx, n := 0, len(vals); idx < n/2; idx++ {
				vals[idx], vals[n-idx-1] = vals[n-idx-1], vals[idx]
			}
		}
		i++

		q = append(q, tmpQ...)
		ans = append(ans, vals)
	}
	return ans
}
