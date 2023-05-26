package simple_and_midium

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var q []*Node
	q = append(q, root)
	for len(q) > 0 {
		width := len(q)
		for i := 0; i < width; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}

			if i < width-1 {
				q[i].Next = q[i+1]
			}
			// i == width - 1 default:
			//q[width-1].Next = nil
		}

		q = q[width:]
	}
	return root
}

func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	// 最下面一层，直接返回
	if root.Left == nil && root.Right == nil {
		return root
	}

	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connect2(root.Left)
	connect2(root.Right)
	return root
}
