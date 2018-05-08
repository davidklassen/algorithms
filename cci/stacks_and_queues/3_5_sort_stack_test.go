package stacks_and_queues

import (
	"testing"
)

func TestSortStack(t *testing.T) {
	s := NewStack()
	s.Push(2)
	s.Push(1)
	s.Push(3)
	SortStack(s)
	if val, _ := s.Pop(); val != 1 {
		t.Error("expected 1")
	}
	if val, _ := s.Pop(); val != 2 {
		t.Error("expected 2")
	}
	if val, _ := s.Pop(); val != 3 {
		t.Error("expected 3")
	}
}
