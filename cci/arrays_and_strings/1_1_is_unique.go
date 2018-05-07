package arrays_and_strings

func IsUnique(s string) bool {
	charset := make(map[rune]struct{})
	for _, ch := range []rune(s) {
		if _, ok := charset[ch]; !ok {
			charset[ch] = struct{}{}
		} else {
			return false
		}
	}
	return true
}
