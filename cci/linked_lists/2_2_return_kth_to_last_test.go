package linked_lists

import (
	"testing"
)

func TestReturnKthToLast(t *testing.T) {
	l := &Node{data: 1, next: &Node{data: 2, next: &Node{data: 3, next: &Node{data: 4}}}}
	if ReturnKthToLast(l, 0).data != 4 {
		t.Error("expected last element")
	}
	if ReturnKthToLast(l, 2).data != 2 {
		t.Error("expected element 2")
	}
	if ReturnKthToLast(l, 3).data != 1 {
		t.Error("expected element 1")
	}
	if ReturnKthToLast(l, 5) != nil {
		t.Error("expected nil")
	}
}
