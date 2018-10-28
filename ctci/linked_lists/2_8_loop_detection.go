package linked_lists

func listLen(head *Node) (l int) {
	for l = 0; head != nil; l++ {
		head = head.next
	}
	return
}

func reverse(head *Node) (r *Node) {
	for head != nil {
		next := head.next
		head.next = r
		r = head
		head = next
	}
	return
}

func findIntersection(head1, head2 *Node) *Node {
	// calculate the lengths of both lists
	len1, len2 := listLen(head1), listLen(head2)
	var short, long *Node
	var shortLen, longLen int
	if len1 > len2 {
		short, shortLen = head2, len2
		long, longLen = head1, len1
	} else {
		short, shortLen = head1, len1
		long, longLen = head2, len2
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
	// both short and long point to the intersection
	return short
}

func findLoop(head *Node) *Node {
	// detect a loop and stop on a node in a loop
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}
	return fast
}

func LoopDetection(head *Node) *Node {
	loop := findLoop(head)
	if loop == nil {
		return nil
	}
	// unlink the list after a node in a loop
	before := loop
	after := loop.next
	before.next = nil
	// reverse the head:before part of the original list
	// this way we will get 2 lists that start at "before" and "after" refs
	// the beginning of the loop will be at the intersection node of both lists
	reverse(head)
	res := findIntersection(before, after)
	// restore the original list
	reverse(before)
	before.next = after
	return res
}
