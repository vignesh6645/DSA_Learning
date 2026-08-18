



type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result strings.Builder

	for _, str := range strs {
		result.WriteString(strconv.Itoa(len(str)))
		result.WriteByte('#')
		result.WriteString(str)
	}

	return result.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}

	i := 0

	for i < len(encoded) {
		j := i

		// Find the '#'
		for encoded[j] != '#' {
			j++
		}

		// Extract the length
		length, _ := strconv.Atoi(encoded[i:j])

		// Move past '#'
		j++

		// Extract exactly `length` characters
		result = append(result, encoded[j:j+length])

		// Move to the next encoded string
		i = j + length
	}

	return result
}