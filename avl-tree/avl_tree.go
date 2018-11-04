package avl_tree

type Node struct {
	val    int
	left   *Node
	right  *Node
	height int
}

type AVLTree struct {
	root *Node
}

func (tree AVLTree) Insert(val int) {
	if tree.root == nil {
		tree.root = &Node{val: val, height: 0}
		return
	}
	insert(tree.root, val)
}

func (tree AVLTree) Delete(val int) {

}

func (tree AVLTree) Contains(val int) bool {
	if tree.root == nil {
		return false
	}
	return contains(tree.root, val)
}

func insert(n *Node, val int) *Node {
	// insert BST node
	if val < n.val {
		if n.left == nil {
			n.left = &Node{val: val, height: 0}
		} else {
			n.left = insert(n.left, val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{val: val, height: 0}
		} else {
			n.right = insert(n.right, val)
		}
	}

	// get left and right heights
	var leftHeight, rightHeight int
	if n.left != nil {
		leftHeight = n.left.height
	}
	if n.right != nil {
		rightHeight = n.right.height
	}

	// calculate balance factor
	bFactor := rightHeight - leftHeight

	if bFactor < -1 {
		if val >= n.left.val {
			// left right case
			n.left = rotateLeft(n.left)
		}
		// left left case
		n = rotateRight(n)
	} else if bFactor > 1 {
		if val < n.right.val {
			// right left case
			n.right = rotateRight(n.right)
		}
		// right right case
		n = rotateLeft(n)
	}

	// update node height
	n.height = max(leftHeight, rightHeight) + 1
	return n
}

func contains(n *Node, val int) bool {
	if n.val == val {
		return true
	}
	if val < n.val {
		if n.left == nil {
			return false
		}
		return contains(n.left, val)
	}
	if n.right == nil {
		return false
	}
	return contains(n.right, val)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rotateRight(n *Node) *Node {
	ret := n.left
	n.left = ret.right
	leftHeight, rightHeight := 0, 0
	if n.left != nil {
		leftHeight = n.left.height
	}
	if n.right != nil {
		rightHeight = n.right.height
	}
	n.height = max(leftHeight, rightHeight) + 1
	ret.right = n
	leftHeight, rightHeight = 0, 0
	if ret.left != nil {
		leftHeight = ret.left.height
	}
	if ret.right != nil {
		rightHeight = ret.right.height
	}
	ret.height = max(leftHeight, rightHeight) + 1
	return ret
}

func rotateLeft(n *Node) *Node {
	ret := n.right
	n.right = ret.left
	leftHeight, rightHeight := 0, 0
	if n.left != nil {
		leftHeight = n.left.height
	}
	if n.right != nil {
		rightHeight = n.right.height
	}
	n.height = max(leftHeight, rightHeight) + 1
	ret.left = n
	leftHeight, rightHeight = 0, 0
	if ret.left != nil {
		leftHeight = ret.left.height
	}
	if ret.right != nil {
		rightHeight = ret.right.height
	}
	ret.height = max(leftHeight, rightHeight) + 1
	return ret
}

func PreOrderTraversal(node *Node, visit func(*Node)) {
	if node != nil {
		visit(node)
		PreOrderTraversal(node.left, visit)
		PreOrderTraversal(node.right, visit)
	}
}
