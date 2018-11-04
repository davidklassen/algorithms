package problems

// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.
//
// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).
//
// The matching should cover the entire input string (not partial).
//
// Examples:
//
// Input:
// s = "aa"
// p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".
//
// Input:
// s = "adceb"
// p = "*a*b"
// Output: true
// Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".
//
// Input:
// s = "acdcb"
// p = "a*c?b"
// Output: false
//
// Note:
//
// s could be empty and contains only lowercase letters a-z.
// p could be empty and contains only lowercase letters a-z, and characters like ? or *.
func Match(s string, p string) bool {
	sBytes := []byte(s)
	pBytes := []byte(p)
	sMax := len(sBytes) - 1
	pMax := len(pBytes) - 1
	sPtr := 0

outer:	for pPtr := 0; pPtr <= pMax; pPtr++ {
		pChar := pBytes[pPtr]
		switch pChar {
		case '?':
			if sPtr > sMax {
				return false
			}
			sPtr++
		case '*':
			if pPtr+1 > pMax {
				return true
			}
			next := pBytes[pPtr+1]
			for {
				if next == '*' {
					continue outer
				}
				if sPtr > sMax {
					return false
				}
				if sBytes[sPtr] == next || next == '?' {
					if Match(string(sBytes[sPtr:]), string(pBytes[pPtr+1:])) {
						return true
					}
				}
				sPtr++
			}
			return false
		default:
			if sPtr > sMax {
				return false
			}
			if sBytes[sPtr] != pChar {
				return false
			}
			sPtr++
		}
	}
	return sPtr == sMax+1
}
