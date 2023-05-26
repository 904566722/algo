package bst

import "fmt"

type treeNode struct {
	val    int
	left   *treeNode
	right  *treeNode
}

type binarySearchTree struct {
	root *treeNode
}

// delete 使用「嫁接」的方式来删除
func (bst *binarySearchTree) delete(val int)  {
	tgt := bst.root
	if tgt == nil {
		return
	}

	var parent *treeNode = nil
	for tgt != nil {
		// 找到目标节点（待删除节点）
		if tgt.val == val {
			break
		}

		parent = tgt
		if val < tgt.val {
			tgt = tgt.left
		} else {
			tgt = tgt.right
		}
	}

	// 不存在
	if tgt == nil {
		return
	}

	// 不存在子树或者只有一棵子树的情况
	// - 若存在子树，则使用子树代替待删除节点即可
	// - 否则直接将待删除节点删除
	if tgt.left == nil || tgt.right == nil {
		if tgt.left == nil {
			tgt = tgt.left
		} else {
			tgt = tgt.right
		}

	// 存在两棵子树的情况
	// child 为右节点根节点
	} else {
		tmp := tgt.right
		// 将待删除节点的左子树嫁接到右子树的最左节点
		for tmp.left != nil {
			tmp = tmp.left
		}
		tmp.left = tgt.left

		tgt = tgt.right
	}

	if parent == nil {
		bst.root = tgt
		return
	}
	if parent.left != nil && parent.left.val == val {
		parent.left = tgt
	} else {
		parent.right = tgt
	}
}

/* 删除节点 */
func (bst *binarySearchTree) remove2(num int) {
	cur := bst.root
	// 若树为空，直接提前返回
	if cur == nil {
		return
	}
	// 待删除节点之前的节点位置
	var pre *treeNode = nil
	// 循环查找，越过叶节点后跳出
	for cur != nil {
		if cur.val == num {
			break
		}
		pre = cur
		if cur.val < num {
			// 待删除节点在右子树中
			cur = cur.right
		} else {
			// 待删除节点在左子树中
			cur = cur.right
		}
	}
	// 若无待删除节点，则直接返回
	if cur == nil {
		return
	}
	// 子节点数为 0 或 1
	if cur.left == nil || cur.right == nil {
		var child *treeNode = nil
		// 取出待删除节点的子节点
		if cur.left != nil {
			child = cur.left
		} else {
			child = cur.right
		}
		// 将子节点替换为待删除节点
		if pre.left == cur {
			pre.left = child
		} else {
			pre.right = child
		}
		// 子节点数为 2
	} else {
		// 获取中序遍历中待删除节点 cur 的下一个节点
		tmp := cur.right
		for tmp.left != nil {
			tmp = tmp.left
		}
		// 递归删除节点 tmp
		bst.remove2(tmp.val)
		// 用 tmp 覆盖 cur
		cur.val = tmp.val
	}
}

func TestBSTDelete() {
	bst := &binarySearchTree{
		root: &treeNode{
			val: 3,
			left: &treeNode{val: 1},
			right: &treeNode{val: 4},
		},
	}
	bst.delete(4)
	fmt.Println(bst.root.val)
}

func TestBSTDelete2() {
	bst := &binarySearchTree{
		root: &treeNode{
			val: 3,
			left: &treeNode{val: 1},
			right: &treeNode{val: 4},
		},
	}
	bst.remove2(3)
	fmt.Println(bst.root.val)
}