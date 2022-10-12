package main

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	max, j1 := 0, 0

	for i := 0; i < len(s2); i++ {
		if s1[i] != s2[i] {
			max++
			if max == 1 {
				j1 = i
			} else if max == 2 {
				//if s1[j1] != s2[i] || s2[j1] != s1[i] {
				//	return false
				//}
				sRunes := []rune(s2)
				t := s2[j1]
				sRunes[j1] = sRunes[i]
				sRunes[i] = rune(t)
				if s1 == string(sRunes) {
					return true
				} else {
					return false
				}
			}
		}
	}
	return false
}
