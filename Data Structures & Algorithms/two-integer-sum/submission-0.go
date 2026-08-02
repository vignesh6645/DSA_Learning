func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int, len(nums))
	for ind, n := range nums {
		numsMap[n] = ind
	}

	for ind, num := range nums {
		diff := target - num
		if val, ok := numsMap[diff]; ok && val != ind {
			return []int{ind, val}
		}
	}
	return []int{}
}
