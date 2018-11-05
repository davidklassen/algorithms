package avl_tree

type Node struct {
	val    int
	left   *Node
	right  *Node
	height int
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

	// calculate balance factor
	bFactor := height(n.right) - height(n.left)

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
	n.height = max(height(n.left), height(n.right)) + 1
	return n
}

func rotateRight(n *Node) *Node {
	ret := n.left
	n.left = ret.right
	n.height = max(height(n.left), height(n.right)) + 1
	ret.right = n
	ret.height = max(height(ret.left), height(ret.right)) + 1
	return ret
}

func rotateLeft(n *Node) *Node {
	ret := n.right
	n.right = ret.left
	n.height = max(height(n.left), height(n.right)) + 1
	ret.left = n
	ret.height = max(height(ret.left), height(ret.right)) + 1
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}
