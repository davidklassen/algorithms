package main

import (
	"fmt"
)

type Node struct {
	val      int
	children []*Node
}

func dfs(n *Node, visit func(*Node), visited map[*Node]struct{}) {
	if _, ok := visited[n]; ok {
		return
	}
	visit(n)
	visited[n] = struct{}{}
	for _, n := range n.children {
		dfs(n, visit, visited)
	}
}

func DFS(n *Node, visit func(*Node)) {
	visited := make(map[*Node]struct{})
	dfs(n, visit, visited)
}

func main() {
	fmt.Println("Hello, DFS!")
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
	DFS(n0, func(n *Node) {
		fmt.Println("Node", n.val)
	})
}
