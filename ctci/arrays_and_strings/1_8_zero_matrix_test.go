package arrays_and_strings

import (
	"testing"
)

func compareZeroMatrix(m1, m2 [][]int, w, h int) bool {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if m1[y][x] != m2[y][x] {
				return false
			}
		}
	}
	return true
}

func TestZeroMatrix(t *testing.T) {
	w := 5
	h := 6
	mx := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}
	expect := [][]int{
		{1, 0, 1, 0, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
	}
	if !compareZeroMatrix(ZeroMatrix(mx, w, h), expect, w, h) {
		t.Error("expected mx and expect to be equal")
	}
}
