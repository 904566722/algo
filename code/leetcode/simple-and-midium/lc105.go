package simple_and_midium

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootIdx1 := 0
	rootIdx2 := indexOf(inorder, preorder[rootIdx1])
	root := &TreeNode{
		Val: preorder[rootIdx1],
	}

	inorder1 := inorder[0:rootIdx2]
	inorder2 := inorder[rootIdx2+1:]
	preorder1 := preorder[1 : 1+len(inorder1)]
	preorder2 := preorder[1+len(inorder1):]
	left := buildTree(preorder1, inorder1)
	right := buildTree(preorder2, inorder2)
	root.Left = left
	root.Right = right
	return root
}

func indexOf(nums []int, tgt int) int {
	for i, num := range nums {
		if tgt == num {
			return i
		}
	}
	return -1
}
