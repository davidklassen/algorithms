package arrays_and_strings

import (
	"testing"
)

func TestOneAway(t *testing.T) {
	if !OneAway("pale", "ple") {
		t.Error("expected pale to be one away from ple")
	}
	if !OneAway("pales", "pale") {
		t.Error("expected pales to be one away from pale")
	}
	if !OneAway("pale", "bale") {
		t.Error("expected pale to be one away from bale")
	}
	if OneAway("pale", "bake") {
		t.Error("expected pale to be more than one away from bake")
	}
}
