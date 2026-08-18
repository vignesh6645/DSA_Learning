func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	for _, word := range strs {
		count := [26]int{}
		for i := 0; i < len(word); i++ {
			count[word[i]-'a']++
		}
		groups[count] = append(groups[count], word)
	}
	result := [][]string{}
	for _, words := range groups {
		result = append(result, words)
	}
	return result
}
