package simple_and_midium



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	valAppeared := map[int]bool{}

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		valAppeared[root.Val] = true
		dfs(root.Left)
		dfs(root.Right)
		return
	}
	dfs(root)

	for val, _ := range valAppeared {
		if valAppeared[k-val] && k-val != val {
			return true
		}
	}
	return false
}