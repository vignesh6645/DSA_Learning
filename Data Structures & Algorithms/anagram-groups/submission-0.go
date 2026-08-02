import "slices"
func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)

	for _, word := range strs {
		chars := []byte(word)
		slices.Sort(chars)
		key := string(chars)
		group[key] = append(group[key], word)
	}
	result := [][]string{}
	for _, val := range group {
		result = append(result, val)
	}

	return result
}
