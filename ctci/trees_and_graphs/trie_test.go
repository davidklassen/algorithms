package trees_and_graphs

import (
	"testing"
)

func TestTrieNode(t *testing.T) {
	tr := NewTrie()
	if !tr.IsEmpty() {
		t.Error("new trie should be empty")
	}
	tr.Insert("")
	if tr.IsEmpty() {
		t.Error("trie with empty string should not be empty")
	}
	if !tr.Contains("") {
		t.Error("trie should contain empty string")
	}
	if tr.Contains("foo") {
		t.Error("trie should not contain foo string")
	}
	tr.Insert("foo")
	if !tr.Contains("foo") {
		t.Error("trie should contain foo string")
	}
}
