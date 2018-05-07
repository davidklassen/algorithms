package arrays_and_strings

import (
	"testing"
)

func TestPalindromePermutation(t *testing.T) {
	in := "Tact Coa"
	if !PalindromePermutation(in) {
		t.Errorf("expected %s to be a palindrome permutation", in)
	}
}
