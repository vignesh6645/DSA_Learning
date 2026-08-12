func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool, len(nums))
	for _, num := range nums {
		numSet[num] = true
	}

	maxSeqCount := 0
	for _, num := range nums {
		if !numSet[num - 1] {
			currentNum := num
			currentCount := 1

			for numSet[currentNum+1] {
				currentNum++
				currentCount++
			}

			if currentCount > maxSeqCount {
				maxSeqCount = currentCount
			}
		}
	}

	return maxSeqCount
}