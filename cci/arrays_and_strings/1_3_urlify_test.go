package arrays_and_strings

import (
	"testing"
)

func TestURLify(t *testing.T) {
	in := "Mr John Smith    "
	expect := "Mr%20John%20Smith"
	got := string(URLify([]byte(in), 13))
	if got != expect {
		t.Errorf("expected %s, got %s", expect, got)
	}
}
