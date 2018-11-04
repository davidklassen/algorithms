package bin_tree

import (
	"fmt"
	"testing"
)

func TestRotateRight(t *testing.T) {
	tree := &Node{
		val: 10,
		left: &Node{
			val: 5,
			left: &Node{
				val:  4,
				left: &Node{val: 3},
			},
			right: &Node{val: 8},
		},
		right: &Node{val: 11},
	}

	expect := &Node{
		val: 5,
		left: &Node{
			val:  4,
			left: &Node{val: 3},
		},
		right: &Node{
			val:   10,
			left:  &Node{val: 8},
			right: &Node{val: 11},
		},
	}

	PreOrderTraversal(expect, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()
	tree = RotateRight(tree)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
}

func TestRotateLeft(t *testing.T) {
	tree := &Node{
		val:  30,
		left: &Node{val: 5},
		right: &Node{
			val:  35,
			left: &Node{val: 32},
			right: &Node{
				val:   40,
				right: &Node{val: 45},
			},
		},
	}

	expect := &Node{
		val: 35,
		left: &Node{
			val:   30,
			left:  &Node{val: 5},
			right: &Node{val: 32},
		},
		right: &Node{
			val:   40,
			right: &Node{val: 45},
		},
	}

	PreOrderTraversal(expect, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()
	tree = RotateLeft(tree)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
}
