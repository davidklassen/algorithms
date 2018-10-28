package trees_and_graphs

type BinaryTreeNode struct {
	val   int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (bt *BinaryTreeNode) IsLeaf() bool {
	return bt.left == nil && bt.right == nil
}
