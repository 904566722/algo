package simple_and_midium

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}
	if len(preorder) == 1 {
		return root
	}

	// 从后序数组中找到左子树
	lRootVal := preorder[1]
	lTreeLen := indexOf(postorder, lRootVal) + 1
	lPreorder := preorder[1 : 1+lTreeLen]
	rPreorder := preorder[1+lTreeLen:]
	lPostorder := postorder[:lTreeLen]
	rPostorder := postorder[lTreeLen : len(postorder)-1]
	lTree := constructFromPrePost(lPreorder, lPostorder)
	rTree := constructFromPrePost(rPreorder, rPostorder)

	root.Left = lTree
	root.Right = rTree
	return root
}
