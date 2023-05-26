package avl

type treeNode struct {
	val    int
	height int	// 节点高度
	left   *treeNode
	right  *treeNode
}

type avlTree struct {
	root *treeNode
}

func (t *avlTree) getHeight(node *treeNode) int {
	// 空节点高度为 -1
	if node == nil {
		return -1
	}
	return node.height
}

func (t *avlTree) updateHeight(node *treeNode) {
	// 节点高度等于最高子树的高度 + 1
	node.height = max(t.getHeight(node.left), t.getHeight(node.right)) + 1
}

// balanceFactor 获取平衡因子
// 设平衡因子为 f，平衡二叉树的平衡因子需满足：-1 <= f <= 1
func (t *avlTree) balanceFactor(node *treeNode) int {
	// 空节点的平衡因子为 0
	if node == nil {
		return 0
	}
	// 节点平衡因子 = 左子树高度 - 右子树高度
	return t.getHeight(node.left) - t.getHeight(node.right)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}


/* Rotation */

// 右旋，返回平衡子树的根节点
func (t *avlTree) rightRotate(node *treeNode) *treeNode {
	child := node.left
	// 消除碰撞，执行旋转
	node.left = child.right
	child.right = node
	// 更新节点高度
	t.updateHeight(node)
	t.updateHeight(child)

	return child
}

//		   10
//		   / \
//        5   15 --node
//				\
//			    20 -- child
//               \
//                30
//
//		   10
//		   / \
//        5   20 --child
//			 /  \
//	node -- 15   30
//
func (t *avlTree) leftRotate(node *treeNode) *treeNode  {
	child := node.right
	// 消除碰撞，执行旋转
	node.right = child.left
	child.left = node
	// 更新节点高度
	t.updateHeight(node)
	t.updateHeight(child)

	return child
}

func (t *avlTree) rotate(node *treeNode) *treeNode {
	// 失衡节点的平衡因子
	bf := t.balanceFactor(node)
	if bf > 1 {
		if t.balanceFactor(node.left) >= 0 {
			return t.rightRotate(node)
		} else {
			// 先左旋，后右旋
			node.left = t.leftRotate(node.left)
			return t.rightRotate(node)
		}
	} else if bf < -1 {
		if t.balanceFactor(node.right) <= 0 {
			return t.leftRotate(node)
		} else {
			// 先右旋，后左旋
			node.right = t.rightRotate(node.right)
			return t.leftRotate(node)
		}
	}
	// 已经是平衡状态
	return node
}

func (t *avlTree) insert(val int)  {
	t.root = t.insertHelper(t.root, val)
}

func (t *avlTree) insertHelper(node *treeNode, val int) *treeNode {
	if node == nil {
		return &treeNode{
			val: val,
		}
	}

	/* 1. 插入 */
	if val < node.val {
		node.left = t.insertHelper(node.left, val)
	} else if val > node.val {
		node.right = t.insertHelper(node.right, val)
	} else {
		// 重复节点不插入
		return node
	}

	// 更新节点高度
	t.updateHeight(node)
	/* 2.执行旋转，恢复平衡 */
	node = t.rotate(node)
	return node
}

func (t *avlTree) remove(val int)  {
	t.removeHelper(t.root, val)
}

func (t *avlTree) removeHelper(node *treeNode, val int) *treeNode {
	if node == nil {
		return nil
	}

	if val < node.val {
		node.left = t.removeHelper(node.left, val)
	} else if val > node.val {
		node.right = t.removeHelper(node.right, val)
	} else {
		if node.left == nil || node.right == nil {
			child := node.left
			if node.right != nil {
				child = node.right
			}

			if child == nil {
				// 没有子树，直接删除节点即可
				return nil
			} else {
				// 有至多一个节点，返回该节点
				return child
			}
		} else {
			// 子节点的数量为 2，删除当前节点的后继节点，并用后继节点的值替换当前节点
			tmp := node.right
			for tmp.left != nil {
				tmp = tmp.left
			}

			// 递归删除这个后继节点
			node.right = t.removeHelper(node.right, tmp.val)
			// 替换当前节点
			node.val = tmp.val
		}
	}

	// 更新节点高度
	t.updateHeight(node)
	/* 执行旋转，保持平衡 */
	node = t.rotate(node)
	return node
}

