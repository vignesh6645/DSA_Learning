func isPalindrome(s string) bool {
	var cleaned strings.Builder
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleaned.WriteRune(unicode.ToLower(ch))
		}
	}
	str := cleaned.String()
	left, right := 0, len(str)-1
	for left < right {
		if str[left] == str[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}