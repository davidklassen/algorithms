package trees_and_graphs

import (
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap()
	if !h.IsEmpty() {
		t.Error("expected a new heap to be empty")
	}
	if _, err := h.Peek(); err == nil {
		t.Error("expected an error from empty heap")
	}
	if _, err := h.Pop(); err == nil {
		t.Error("expected an error from empty heap")
	}
	h.Push(1)
	if m, _ := h.Peek(); m != 1 {
		t.Error("expected to have 1 on top")
	}
	if m, _ := h.Pop(); m != 1 {
		t.Error("expected to have 1 on top")
	}
	h.Push(1)
	h.Push(2)
	if m, _ := h.Pop(); m != 1 {
		t.Error("expected to have 1 on top")
	}
	if m, _ := h.Pop(); m != 2 {
		t.Error("expected to have 2 on top")
	}
}
