package arrays_and_strings

func checkInsert(s1, s2 string) bool {
	edited := false
	for i, j := 0, 0; i < len(s1); i, j = i+1, j+1 {
		if s1[i] != s2[j] {
			if edited {
				return false
			}
			edited = true
			i--
		}
	}
	return true
}

func checkReplace(s1, s2 string) bool {
	edited := false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if edited {
				return false
			}
			edited = true
		}
	}
	return true
}

func OneAway(s1, s2 string) bool {
	if len(s1) > len(s2) {
		tmp := s1
		s1 = s2
		s2 = tmp
	}
	diff := len(s2) - len(s1)
	if diff > 1 {
		return false
	}
	if diff == 1 {
		return checkInsert(s1, s2)
	}
	return checkReplace(s1, s2)
}
