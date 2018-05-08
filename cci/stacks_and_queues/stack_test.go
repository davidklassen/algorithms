package stacks_and_queues

import (
	"testing"
)

func TestStack(t *testing.T) {
	one := 1
	s := NewStack()
	if !s.IsEmpty() {
		t.Error("new stack should be empty")
	}
	s.Push(one)
	if s.IsEmpty() {
		t.Error("expected to have one element")
	}
	if res, err := s.Peek(); err != nil {
		t.Error(err)
	} else if res != one {
		t.Error("expected to have one on top")
	}
	if res, err := s.Pop(); err != nil {
		t.Error(err)
	} else if res != one {
		t.Error("expected to have one on top")
	}
	if !s.IsEmpty() {
		t.Error("expected to be be empty")
	}
	if _, err := s.Peek(); err == nil {
		t.Error("expected an error from empty stack")
	}
	if _, err := s.Pop(); err == nil {
		t.Error("expected an error from empty stack")
	}
}
