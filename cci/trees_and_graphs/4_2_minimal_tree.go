package trees_and_graphs

func MinimalTree(arr []int) *BinaryTreeNode {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return &BinaryTreeNode{val: arr[0]}
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid+1:]
	return &BinaryTreeNode{
		val:   arr[mid],
		left:  MinimalTree(left),
		right: MinimalTree(right),
	}
}
