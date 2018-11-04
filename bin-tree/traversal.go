package bin_tree

func InOrderTraversal(node *Node, visit func(*Node)) {
	if node != nil {
		InOrderTraversal(node.left, visit)
		visit(node)
		InOrderTraversal(node.right, visit)
	}
}

func PreOrderTraversal(node *Node, visit func(*Node)) {
	if node != nil {
		visit(node)
		PreOrderTraversal(node.left, visit)
		PreOrderTraversal(node.right, visit)
	}
}

func PostOrderTraversal(node *Node, visit func(*Node)) {
	if node != nil {
		PostOrderTraversal(node.left, visit)
		PostOrderTraversal(node.right, visit)
		visit(node)
	}
}
