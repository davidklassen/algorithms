package trees_and_graphs

func InOrderTraversal(node *BinaryTreeNode, visit func(*BinaryTreeNode)) {
	if node != nil {
		InOrderTraversal(node.left, visit)
		visit(node)
		InOrderTraversal(node.right, visit)
	}
}

func PreOrderTraversal(node *BinaryTreeNode, visit func(*BinaryTreeNode)) {
	if node != nil {
		visit(node)
		PreOrderTraversal(node.left, visit)
		PreOrderTraversal(node.right, visit)
	}
}

func PostOrderTraversal(node *BinaryTreeNode, visit func(*BinaryTreeNode)) {
	if node != nil {
		PostOrderTraversal(node.left, visit)
		PostOrderTraversal(node.right, visit)
		visit(node)
	}
}
