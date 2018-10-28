package arrays_and_strings

import (
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	a := "abcd"
	b := "bcda"
	c := "foobar"
	if !CheckPermutation(a, b) {
		t.Errorf("expected %s to be a permutation of %s", a, b)
	}
	if CheckPermutation(a, c) {
		t.Errorf("expected %s not to be a permutation of %s", a, c)
	}
}
