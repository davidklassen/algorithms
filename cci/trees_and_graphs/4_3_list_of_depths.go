package trees_and_graphs

func traverse(n *BinaryTreeNode, depth int, depths [][]*BinaryTreeNode) [][]*BinaryTreeNode {
	if n == nil {
		return depths
	}
	if depth > len(depths) - 1 {
		depths = append(depths, nil)
	}
	depths[depth] = append(depths[depth], n)
	depths = traverse(n.left, depth+1, depths)
	depths = traverse(n.right, depth+1, depths)
	return depths
}

func ListOfDepths(n *BinaryTreeNode) [][]*BinaryTreeNode {
	return traverse(n, 0, nil)
}
