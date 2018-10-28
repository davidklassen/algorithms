package linked_lists

import (
	"testing"
)

func TestLoopDetection(t *testing.T) {
	last := &Node{data: 8}
	tail := &Node{data: 4, next: &Node{data: 5, next: &Node{data: 6, next: &Node{data: 7, next: last}}}}
	last.next = tail
	l := &Node{data: 1, next: &Node{data: 2, next: &Node{data: 3, next: tail}}}
	if LoopDetection(l).data != 4 {
		t.Error("expected to find a loop at node 4")
	}
	noLoop := &Node{data: 1, next: &Node{data: 2}}
	if LoopDetection(noLoop) != nil {
		t.Error("expected to have no loops")
	}
}
