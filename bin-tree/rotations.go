package bin_tree

func RotateRight(n *Node) *Node {
	ret := n.left
	n.left = ret.right
	ret.right = n
	return ret
}

func RotateLeft(n *Node) *Node {
	ret := n.right
	n.right = ret.left
	ret.left = n
	return ret
}
