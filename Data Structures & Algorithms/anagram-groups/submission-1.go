func groupAnagrams(strs []string) [][]string {
	group := make(map[[26]int][]string)

	for _, word := range strs {
		var count [26]int
		for i := 0; i < len(word); i++ {
			count[word[i]-'a']++
		}
		group[count] = append(group[count], word)
	}
	result := [][]string{}
	for _, val := range group {
		result = append(result, val)
	}

	return result
}