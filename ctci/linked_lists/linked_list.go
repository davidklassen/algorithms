package linked_lists

type Node struct {
	next *Node
	data int
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Add(data int) {
	ll.head = &Node{
		next: ll.head,
		data: data,
	}
}

func (ll *LinkedList) Find(data int) *Node {
	current := ll.head
	for current != nil {
		if current.data == data {
			return current
		}
	}
	return nil
}

func (ll *LinkedList) Delete(node *Node) {
	if node == ll.head {
		ll.head = ll.head.next
		return
	}
	if node.next == nil {
		node = node.next
		return
	}
	*node = *node.next
}

func (ll *LinkedList) Head() *Node {
	return ll.head
}
