package arrays_and_strings

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	empty := ""
	unique := "abcdefg"
	nonUnique := "aabcdefg"
	if !IsUnique(empty) {
		t.Errorf("expected empty string to be unique")
	}
	if !IsUnique(unique) {
		t.Errorf("expected %s to be unique", unique)
	}
	if IsUnique(nonUnique) {
		t.Errorf("expected %s not to be unique", nonUnique)
	}
}
