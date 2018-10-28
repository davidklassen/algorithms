package linked_lists

func RemoveDups(head *Node) *Node {
	if head == nil {
		return nil
	}
	counts := make(map[int]struct{})
	counts[head.data] = struct{}{}
	prev := head
	cur := head.next
	for cur != nil {
		if _, ok := counts[cur.data]; !ok {
			counts[cur.data] = struct{}{}
		} else {
			prev.next = cur.next
		}
		prev = prev.next
		cur = cur.next
	}
	return head
}
