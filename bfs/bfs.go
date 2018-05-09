package main

import (
	"errors"
	"fmt"
)

type Node struct {
	val      int
	children []*Node
}

type Queue struct {
	nodes []*Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(n *Node) {
	q.nodes = append(q.nodes, n)
}

func (q *Queue) Dequeue() (n *Node, err error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	n, q.nodes = q.nodes[0], q.nodes[1:]
	return n, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}

func BFS(n *Node, visit func(*Node)) {
	q := NewQueue()
	q.Enqueue(n)
	visited := make(map[*Node]struct{})
	for !q.IsEmpty() {
		cur, _ := q.Dequeue()
		visit(cur)
		visited[cur] = struct{}{}
		for _, n := range cur.children {
			if _, ok := visited[n]; !ok {
				q.Enqueue(n)
			}
		}
	}
}

func main() {
	fmt.Println("Hello, BFS!")
	n0 := &Node{val: 0}
	n1 := &Node{val: 1}
	n2 := &Node{val: 2}
	n3 := &Node{val: 3}
	n4 := &Node{val: 4}
	n5 := &Node{val: 5}
	n0.children = []*Node{n1, n4, n5}
	n1.children = []*Node{n3, n4}
	n2.children = []*Node{n1}
	n3.children = []*Node{n2, n4}
	BFS(n0, func(n *Node) {
		fmt.Println("Node", n.val)
	})
}
