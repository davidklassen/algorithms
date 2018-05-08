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

func findIntersection(head1, head2 *Node) *Node {
	// calculate the lengths of both lists
	len1 := listLen(head1)
	len2 := listLen(head2)
	var short, long *Node
	var shortLen, longLen int
	if len1 > len2 {
		short = head2
		shortLen = len2
		long = head1
		longLen = len1
	} else {
		short = head1
		shortLen = len1
		long = head2
		longLen = len2
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
