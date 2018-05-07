package arrays_and_strings

func CheckPermutation(s1, s2 string) bool {
	counters := make(map[rune]int)
	if len(s1) != len(s2) {
		return false
	}
	for _, ch := range []rune(s1) {
		counters[ch] += 1
	}
	for _, ch := range []rune(s2) {
		counters[ch] -= 1
		if counters[ch] < 0 {
			return false
		}
	}
	return true
}
