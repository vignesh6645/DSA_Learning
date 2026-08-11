func topKFrequent(nums []int, k int) []int {
	// Count frequency of each number
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	// bucket[i] contains numbers that appear i times
	buckets := make([][]int, len(nums)+1)

	for num, count := range countMap {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)

	// Start from highest frequency
	for freq := len(buckets) - 1; freq >= 0 && len(result) < k; freq-- {
		for _, num := range buckets[freq] {
			result = append(result, num)

			if len(result) == k {
				break
			}
		}
	}

	return result
}
