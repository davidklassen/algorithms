package arrays_and_strings

import (
	"strconv"
)

func StringCompression(s string) string {
	if s == "" {
		return s
	}
	compressed := []byte(s)[:1]
	currentChar := compressed[0]
	currentCount := 1
	for i := 1; i < len(s); i++ {
		if s[i] == currentChar {
			currentCount++
		} else {
			currentChar = s[i]
			compressed = append(compressed, []byte(strconv.Itoa(currentCount))...)
			compressed = append(compressed, currentChar)
			currentCount = 1
		}
	}
	compressed = append(compressed, []byte(strconv.Itoa(currentCount))...)
	if len(compressed) >= len(s) {
		return s
	}
	return string(compressed)
}
