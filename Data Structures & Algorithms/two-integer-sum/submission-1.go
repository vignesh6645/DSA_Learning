func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))

	for ind, num := range nums {
		diff := target - num
		if i, ok := numsMap[diff]; ok {
			return []int{i, ind}
		}
		numsMap[num] = ind
	}
	return []int{}
}
