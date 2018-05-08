package linked_lists

func listLen(head *Node) int {
	l := 0
	for head != nil {
		l++
		head = head.next
	}
	return l
}

func reverse(head *Node) *Node {
	var r *Node = nil
	for head != nil {
		next := head.next
		head.next = r
		r = head
		head = next
	}
	return r
}

func LoopDetection(head *Node) *Node {
	// detect a loop and stop on a node in a loop
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}

	// return nil if no loop detected
	if fast == nil {
		return nil
	}

	// unlink the list after a node in a loop
	before := fast
	after := fast.next
	before.next = nil

	// reverse the head-before part
	reverse(head)

	// calculate the lengths of both lists
	beforeLen := listLen(before)
	afterLen := listLen(after)

	var short, long *Node
	var shortLen, longLen int
	if beforeLen > afterLen {
		short = after
		shortLen = afterLen
		long = before
		longLen = beforeLen
	} else {
		short = before
		shortLen = beforeLen
		long = after
		longLen = afterLen
	}

	// traverse long list longLen - shortLen times
	for diff := longLen - shortLen; diff > 0; diff-- {
		long = long.next
	}

	// find intersection
	for short != long {
		short = short.next
		long = long.next
	}

	// reverse and link back
	reverse(before)
	before.next = after

	// both short and long point to the loop intersection
	return short
}
