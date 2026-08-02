func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	anagramMap := make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		anagramMap[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		anagramMap[t[i]]--
		if anagramMap[t[i]] < 0 {
			return false
		}
		if anagramMap[t[i]] == 0 {
			delete(anagramMap, t[i])
		}
	}

	return len(anagramMap) == 0
}
