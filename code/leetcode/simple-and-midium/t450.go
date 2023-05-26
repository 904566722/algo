package simple_and_midium

import (
	"fmt"
)

// 如果找到该节点，有四种情况：
//	没有左节点，直接将右子树代替该节点的位置
//	没有右节点，直接将左子树代替该节点的位置
//	均有左右节点，根据中序节点的性质，当前的序列为：左子树、当前节点、右子树的最左节点
//			删除当前节点之后，为了保持中序，需要将左子树嫁接到右子树的最左节点的左节点，然后将右子树的根节点替换当前节点的位置
//	均没有左右节点（叶子节点），直接删除该节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	var curParent *TreeNode = nil
	cur := root
	for cur != nil && cur.Val != key {
		curParent = cur
		if key < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	// not found
	if cur == nil {
		return root
	}

	if cur.Left != nil && cur.Right != nil {
		rightMinNode := cur.Right
		for rightMinNode.Left != nil {
			rightMinNode = rightMinNode.Left
		}

		rightMinNode.Left = cur.Left
		cur = cur.Right
	} else if cur.Left != nil {
		cur = cur.Left
	} else if cur.Right != nil {
		cur = cur.Right
	} else {
		cur = nil
	}

	if curParent == nil {
		return cur
	}
	if curParent.Left != nil && curParent.Left.Val == key {
		curParent.Left = cur
	} else {
		curParent.Right = cur
	}
	return root
}

func testptr() {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	l := root.Left
	l.Val = 10
	l = nil
	fmt.Println(root.Left)
	fmt.Println(l)
}
