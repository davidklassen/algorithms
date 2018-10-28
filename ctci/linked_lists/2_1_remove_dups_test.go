package linked_lists

import (
	"testing"
)

func TestRemoveDups(t *testing.T) {
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

	l := &Node{data: 1, next: &Node{data: 2, next: &Node{data: 1}}}
	expect := &Node{data: 1, next: &Node{data: 2}}
	if !compare(RemoveDups(l), expect) {
		t.Error("expected l to equal expect")
	}
}
