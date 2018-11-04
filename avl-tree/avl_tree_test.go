package avl_tree

import (
	"fmt"
	"testing"
)

func TestAVLTree(t *testing.T) {
	tree := &Node{val: 1}
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()

	tree = insert(tree, 2)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()

	tree = insert(tree, 3)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()

	tree = insert(tree, 4)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()

	tree = insert(tree, 5)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()

	tree = insert(tree, 6)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()

	tree = insert(tree, 7)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()

	tree = insert(tree, 8)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()

	tree = insert(tree, 9)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()

	tree = insert(tree, 10)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()

	tree = insert(tree, 11)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()

	tree = insert(tree, 12)
	PreOrderTraversal(tree, func(n *Node) {
		fmt.Printf("%d ", n.val)
	})
	fmt.Println()
}
