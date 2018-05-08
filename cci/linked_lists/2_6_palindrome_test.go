package linked_lists

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	l := &Node{data: 1, next: &Node{data: 2, next: &Node{data: 3}}}
	p := &Node{data: 1, next: &Node{data: 2, next: &Node{data: 1}}}
	if Palindrome(l) {
		t.Error("l is not a palindrome")
	}
	if !Palindrome(p) {
		t.Error("p is a palindrome")
	}
}
