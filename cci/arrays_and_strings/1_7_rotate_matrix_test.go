package arrays_and_strings

import (
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	in := "abcdefghijklmnopqrstuvwxyzabcdefghij"
	expect := "yzabmnopabcdcdefqrstefghghijuvwxijkl"
	if string(RotateMatrix([]byte(in), 3)) != expect {
		t.Errorf("expected %s to be rotated to %s", in, expect)
	}
}
