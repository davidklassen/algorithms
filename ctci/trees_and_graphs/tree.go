package trees_and_graphs

type TreeNode struct {
	val      int
	children []*TreeNode
}

func (t *TreeNode) IsLeaf() bool {
	return len(t.children) == 0
}

type Tree struct {
	root *TreeNode
}
