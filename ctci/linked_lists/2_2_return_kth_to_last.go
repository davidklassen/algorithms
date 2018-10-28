package linked_lists

func ReturnKthToLast(head *Node, k int) *Node {
	var res *Node = nil
	cur := head
	for i := 0; i < k && cur != nil; i++ {
		cur = cur.next
	}
	if cur == nil {
		return nil
	}
	cur = cur.next
	res = head
	for cur != nil {
		res = res.next
		cur = cur.next
	}
	return res
}
