package bin_tree

import (
	"testing"
)

func TestBST(t *testing.T) {
	var tree BST
	if tree.Contains(1) {
		t.Error("expected tree to not contain '1'")
	}
	tree.Insert(1)
	tree.Insert(1)
	if !tree.Contains(1) {
		t.Error("expected tree to contain '1'")
	}
	tree.Delete(1)
	if !tree.Contains(1) {
		t.Error("expected tree to contain '1'")
	}
	tree.Delete(1)
	if tree.Contains(1) {
		t.Error("expected tree to not contain '1'")
	}
}
