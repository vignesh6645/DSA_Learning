func isAnagram(s string, t string) bool {
	if len(s) != len(t) { // 1 operation
		return false // 1 operation
	}
	anagramMap := make(map[byte]int, len(s)) // 1 operation

	for i := 0; i < len(s); i++ { // n operation
		anagramMap[s[i]]++ // 1 operation
	}

	for i := 0; i < len(t); i++ { // n operation
		anagramMap[t[i]]--        // 1 operation
		if anagramMap[t[i]] < 0 { // 1 operation
			return false // 1 operation
		}
	}

	return true // 1 operation
	// y = 1 + 1 + 1 + 2n + 5n + 1 => O(n)
}