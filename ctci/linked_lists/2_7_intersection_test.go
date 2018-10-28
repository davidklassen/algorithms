package linked_lists

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	tail := &Node{data: 4, next: &Node{data: 5}}
	l1 := &Node{data: 1, next: &Node{data: 2, next: &Node{data: 3, next: tail}}}
	l2 := &Node{data: 1, next: &Node{data: 2, next: tail}}
	l3 := &Node{data: 1, next: &Node{data: 2, next: &Node{data: 3, next: &Node{data: 4}}}}
	if Intersection(l1, l3) != nil {
		t.Error("l1 and l3 should not intersect")
	}
	if Intersection(l1, l2) != tail {
		t.Error("l1 and l2 should intersect")
	}
}
