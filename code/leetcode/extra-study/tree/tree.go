package tree

import (
	"fmt"
)

type treeNode struct {
	val int
	left *treeNode
	right *treeNode
}

// 前序遍历: 访问根节点，以前序遍历的方式访问左节点，以前序遍历的方式访问右节点
func pOrder(root *treeNode)  {
	// end
	if root == nil {
		return
	}

	pOrder(root.left)
	pOrder(root.right)
	return
}

func pOrderStack(root *treeNode)  {
	if root == nil {
		return
	}

	var st []*treeNode
	st = append(st, root)
	var pnt []int
	for len(st) > 0 {
		// pop
		cur := st[len(st)-1]
		st = st[:len(st)-1]
		pnt = append(pnt, cur.val)

		// judge left\right
		if cur.right != nil {
			st = append(st, cur.right)
		}
		if cur.left != nil {
			st = append(st, cur.left)
		}
	}
	fmt.Println(pnt)
}

// 按中序遍历访问左节点，访问节点元素，按中序遍历访问右节点
func ldrOrder(root *treeNode)  {
	if root == nil {
		return
	}

	ldrOrder(root.left)
	fmt.Printf("%d ", root.val)
	ldrOrder(root.right)

	return
}

//
//
//  left
func ldrOrderStack(root *treeNode) {
	if root == nil {
		return
	}

	var st []*treeNode
	cur := root
	// 找到最左叶子节点
	for cur != nil {
		st = append(st, cur)
		cur = cur.left
	}

	// 出栈，访问元素
	// 如果存在右节点，继续入栈，直到最左叶子节点
	var pnt []int
	for len(st) > 0 {
		cur := st[len(st)-1]
		st = st[:len(st)-1]
		pnt = append(pnt, cur.val)

		tmp := cur.right
		for tmp != nil {
			st = append(st, tmp)
			tmp = tmp.left
		}
	}

	return
}

func ldrOrderStack2(root *treeNode) {
	if root == nil {
		return
	}

	var st []*treeNode
	// 出栈，访问元素
	// 如果存在右节点，继续入栈，直到最左叶子节点
	var pnt []int
	tmp := root
	for len(st) > 0 || tmp != nil {
		// 入栈直到最左节点
		for tmp != nil {
			st = append(st, tmp)
			tmp = tmp.left
		}

		// 出栈，访问节点元素
		tmp = st[len(st)-1]
		st = st[:len(st)-1]
		pnt = append(pnt, tmp.val)
		// 继续找右节点的最左叶子节点
		tmp = tmp.right
	}

	return
}

func lrdOrder(root *treeNode)  {
	if root == nil {
		return
	}

	lrdOrder(root.left)
	lrdOrder(root.right)
	fmt.Printf("%d ", root.val)

	return
}

func lrdOrderStack(root *treeNode)  {
	var st []*treeNode
	tmpRoot := root
	pre := root
	for len(st) > 0 || tmpRoot != nil {
		for tmpRoot != nil {
			st = append(st, tmpRoot)
			tmpRoot = tmpRoot.left
		}

		tmpRoot = st[len(st)-1]
		// 当节点右节点为空 或者 右子树已经访问过的情况下，访问根节点
		if tmpRoot.right == nil || tmpRoot.right == pre {
			fmt.Printf("%d ", tmpRoot.val)
			pre = tmpRoot
			tmpRoot = nil	// 将当前节点标记为空，避免下一个循环又将该节点入栈
			st = st[:len(st)-1]
		} else {
			// 右节点还没访问，后序遍历右节点
			tmpRoot = tmpRoot.right
		}
	}
	return
}

func levelOrder(root *treeNode)  {
	if root == nil {
		return
	}
	var queue []*treeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", cur.val)

		if cur.left != nil {
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			queue = append(queue, cur.right)
		}
	}
	return
}