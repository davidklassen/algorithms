package stacks_and_queues

import (
	"errors"
)

type queueNode struct {
	val  int
	next *queueNode
}

type Queue struct {
	first *queueNode
	last  *queueNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val int) {
	node := &queueNode{val: val}
	if q.IsEmpty() {
		q.last = node
		q.first = q.last
	} else {
		q.last.next = node
	}

}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	res := q.first.val
	q.first = q.first.next
	return res, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return q.first.val, nil
}

func (q *Queue) IsEmpty() bool {
	return q.first == nil
}
