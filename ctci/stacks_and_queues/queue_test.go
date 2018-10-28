package stacks_and_queues

import (
	"testing"
)

func TestQueue(t *testing.T) {
	one := 1
	q := NewQueue()
	if !q.IsEmpty() {
		t.Error("new queue should be empty")
	}
	q.Enqueue(one)
	if q.IsEmpty() {
		t.Error("expected to have one element")
	}
	if res, err := q.Peek(); err != nil {
		t.Error(err)
	} else if res != one {
		t.Error("expected to have one on top")
	}
	if res, err := q.Dequeue(); err != nil {
		t.Error(err)
	} else if res != one {
		t.Error("expected to have one on top")
	}
	if !q.IsEmpty() {
		t.Error("expected to be be empty")
	}
	if _, err := q.Peek(); err == nil {
		t.Error("expected an error from empty queue")
	}
	if _, err := q.Dequeue(); err == nil {
		t.Error("expected an error from empty queue")
	}
}
