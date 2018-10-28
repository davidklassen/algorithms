package arrays_and_strings

func URLify(s []byte, l int) []byte {
	end := len(s) - 1
	for i := l - 1; i >= 0; i-- {
		if s[i] != ' ' {
			s[end] = s[i]
			end--
		} else {
			// insert "%20"
			s[end] = '0'
			s[end-1] = '2'
			s[end-2] = '%'
			end = end - 3
		}
	}
	return s
}
