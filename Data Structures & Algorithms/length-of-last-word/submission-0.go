func lengthOfLastWord(s string) int {
	ans := 0
	str := strings.TrimSpace(s)
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			break
		}
		ans++
	}
	return ans
}