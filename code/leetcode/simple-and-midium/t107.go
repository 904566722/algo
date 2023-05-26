package simple_and_midium


func levelOrderBottom(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}

	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		tmpQlen := len(q)
		var vals []int
		for i:=0; i<tmpQlen; i++ {
			vals = append(vals, q[i].Val)

			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[tmpQlen:]
		ans = append(ans, vals)
	}
	for i,n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return ans
}