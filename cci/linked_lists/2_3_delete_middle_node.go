package linked_lists

func DeleteMiddleNode(node *Node) {
	node.data = node.next.data
	node.next = node.next.next
}
