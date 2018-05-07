package linked_lists

func reverse(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	var r *Node = nil
	for cur != nil {
		tmp := cur.next
		cur.next = r
		r = cur
		cur = tmp
	}
	return r
}

func SumLists(l1, l2 *Node) *Node {
	var a, b int
	pos := 1
	cur := l1
	for cur != nil {
		a += pos * cur.data
		pos *= 10
		cur = cur.next
	}
	cur = l2
	pos = 1
	for cur != nil {
		b += pos * cur.data
		pos *= 10
		cur = cur.next
	}

	sum := a + b
	var sumList *Node
	for sum > 0 {
		dig := sum % 10
		sum = sum / 10
		if sumList == nil {
			sumList = &Node{data: dig}
		} else {
			sumList = &Node{
				data: dig,
				next: sumList,
			}
		}
	}
	return reverse(sumList)
}
