package arrays_and_strings

import (
	"testing"
)

func TestStringRotation(t *testing.T) {
	if !StringRotation("waterbottle", "erbottlewat") {
		t.Error("expected waterbottle to be a string rotation of erbottlewat")
	}
}
