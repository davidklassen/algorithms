package arrays_and_strings

import (
	"strings"
	"unicode"
)

func PalindromePermutation(s string) bool {
	counts := make(map[rune]uint)
	for _, r := range []rune(strings.ToLower(s)) {
		if !unicode.IsSpace(r) {
			counts[r] += 1
		}
	}
	odd := 0
	for _, c := range counts {
		if c%2 != 0 {
			odd++
		}
		if odd > 1 {
			return false
		}
	}
	return true
}
