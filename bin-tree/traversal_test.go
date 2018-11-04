package bin_tree

import (
	"testing"
)

var tree = &Node{
	val: 1,
	left: &Node{
		val: 2,
		left: &Node{
			val: 3,
		},
		right: &Node{
			val: 4,
		},
	},
	right: &Node{
		val: 5,
		left: &Node{
			val: 6,
		},
		right: &Node{
			val: 7,
		},
	},
}

func CompareSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func Traversal(t *testing.T, tree *Node, traversal func(*Node, func(*Node)), expect []int) {
	var got []int
	traversal(tree, func(node *Node) {
		got = append(got, node.val)
	})
	if !CompareSlices(got, expect) {
		t.Errorf("expected %#v, got %#v", expect, got)
	}
}

func TestInOrderTraversal(t *testing.T) {
	Traversal(t, tree, InOrderTraversal, []int{3, 2, 4, 1, 6, 5, 7})
}

func TestPreOrderTraversal(t *testing.T) {
	Traversal(t, tree, PreOrderTraversal, []int{1, 2, 3, 4, 5, 6, 7})
}

func TestPostOrderTraversal(t *testing.T) {
	Traversal(t, tree, PostOrderTraversal, []int{3, 4, 2, 6, 7, 5, 1})
}
