package linked_lists

import (
	"testing"
)

func TestSumLists(t *testing.T) {
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

	l1 := &Node{data: 7, next: &Node{data: 1, next: &Node{data: 6}}}
	l2 := &Node{data: 5, next: &Node{data: 9, next: &Node{data: 2}}}
	expect := &Node{data: 2, next: &Node{data: 1, next: &Node{data: 9}}}
	if !compare(SumLists(l1, l2), expect) {
		t.Error("expected to get a sum")
	}
}
