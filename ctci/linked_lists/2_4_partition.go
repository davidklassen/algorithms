package linked_lists

func addNode(head *Node, data int) *Node {
	return &Node{
		data: data,
		next: head,
	}
}

func deleteNode(head *Node, cur *Node) *Node {
	if head == nil {
		return head
	}
	if head == cur {
		return cur.next
	}
	prev := head
	for prev.next != cur {
		prev = prev.next
	}
	prev.next = cur.next
	return head
}

func Partition(head *Node, x int) *Node {
	var less *Node
	cur := head
	for cur != nil {
		if cur.data < x {
			less = addNode(less, cur.data)
			head = deleteNode(head, cur)
		}
		cur = cur.next
	}
	bigger := head
	lessEnd := less
	for lessEnd.next != nil {
		lessEnd = lessEnd.next
	}
	lessEnd.next = bigger
	return less
}
