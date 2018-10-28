package stacks_and_queues

import (
	"errors"
)

type stackNode struct {
	val  int
	next *stackNode
}

type Stack struct {
	head *stackNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	s.head = &stackNode{
		val:  val,
		next: s.head,
	}
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	res := s.head.val
	s.head = s.head.next
	return res, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.head.val, nil
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}
