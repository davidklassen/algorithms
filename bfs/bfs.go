package main

import (
	"fmt"
)

type node struct {
	val   int
	left  *node
	right *node
}

type list struct {
	val  *node
	next *list
	prev *list
}

type queue struct {
	first *list
	last  *list
}

func (q *queue) enqueue(n *node) {
	if n == nil {
		return
	}
	el := &list{
		val:  n,
		next: q.first,
	}
	if q.last == nil {
		q.last = el
	}
	q.first = el
	if q.first.next != nil {
		q.first.next.prev = el
	}
}

func (q *queue) dequeue() *node {
	if q.last == nil {
		return nil
	}
	el := q.last
	if q.first == q.last {
		q.first = nil
		q.last = nil
	} else {
		q.last = q.last.prev
		q.last.next = nil
	}
	return el.val
}

func (q *queue) isEmpty() bool {
	return q.first == nil
}

func visit(n *node) {
	fmt.Println(n.val)
}

func bfs(node *node) {
	queue := queue{}
	queue.enqueue(node)
	for !queue.isEmpty() {
		current := queue.dequeue()
		visit(current)
		queue.enqueue(current.left)
		queue.enqueue(current.right)
	}
}

func main() {
	bfs(&node{
		val: 1,
		left: &node{
			val: 2,
			left: &node{
				val: 3,
				left: &node{
					val: 4,
				},
			},
			right: &node{
				val: 5,
			},
		},
		right: &node{
			val: 6,
			left: &node{
				val: 7,
			},
		},
	})
}
