package linked_lists

import (
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	compare := func(l1, l2 *Node) bool {
		for l1 != nil {
			if l2 == nil {
				return false
			}
			if l1.data != l2.data {
				return false
			}
			l1 = l1.next
			l2 = l2.next
		}
		return l2 == nil
	}

	middle := &Node{data: 2, next: &Node{data: 3, next: &Node{data: 4}}}
	l := &Node{data: 1, next: middle}
	expect := &Node{data: 1, next: &Node{data: 3, next: &Node{data: 4}}}
	DeleteMiddleNode(middle)
	if !compare(l, expect) {
		t.Error("expected the middle node to be removed")
	}
}
