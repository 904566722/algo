package simple_and_midium

// 层序遍历，判断每一层的数是否回文，空节点的数用-1表示
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var q []*TreeNode

	q = append(q, root)
	for len(q) > 0 {

		var nums []int
		// get head
		n := len(q)
		for i := 0; i < n; i++ {
			if q[i] != nil {
				nums = append(nums, q[i].Val)
				q = append(q, q[i].Left)
				q = append(q, q[i].Right)
			} else {
				nums = append(nums, -101)
			}
		}
		q = q[n:]
		// jud
		for i, n := 0, len(nums); i < n/2; i++ {
			if nums[i] != nums[n-i-1] {
				return false
			}
		}
	}
	return true
}
