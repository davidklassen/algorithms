package bin_tree

type BSTNode struct {
	val   int
	count int
	left  *BSTNode
	right *BSTNode
}

type BST struct {
	root *BSTNode
}

func (tree *BST) Insert(val int) {
	tree.root = insert(tree.root, val)
}

func (tree *BST) Delete(val int) {
	tree.root = del(tree.root, val)
}

func (tree *BST) Contains(val int) bool {
	return contains(tree.root, val)
}

func insert(n *BSTNode, val int) *BSTNode {
	if n == nil {
		return &BSTNode{val: val, count: 1}
	} else if val < n.val {
		n.left = insert(n.left, val)
	} else if val > n.val {
		n.right = insert(n.right, val)
	} else {
		n.count++
	}
	return n
}

func contains(n *BSTNode, val int) bool {
	if n == nil {
		return false
	} else if val < n.val {
		return contains(n.left, val)
	} else if val > n.val {
		return contains(n.right, val)
	}
	return true
}

func del(n *BSTNode, val int) *BSTNode {
	if n == nil {
		return nil
	} else if val < n.val {
		n.left = del(n.left, val)
	} else if val > n.val {
		n.right = del(n.right, val)
	} else {
		if n.count > 1 {
			n.count--
			return n
		}
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}
		inSucc := min(n.right)
		n.val = inSucc.val
		n.count = inSucc.count
		inSucc.count = 0
		n.right = del(n.right, inSucc.val)
	}
	return n
}

func min(n *BSTNode) *BSTNode {
	for n.left != nil {
		n = n.left
	}
	return n
}
