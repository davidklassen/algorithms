package arrays_and_strings

import (
	"testing"
)

func TestStringCompression(t *testing.T) {
	if StringCompression("aabcccccaaa") != "a2b1c5a3" {
		t.Error("expected to compress aabcccccaaa to a2b1c5a3")
	}
	if StringCompression("abc") != "abc" {
		t.Error("expected to compress abc to abc")
	}
}
