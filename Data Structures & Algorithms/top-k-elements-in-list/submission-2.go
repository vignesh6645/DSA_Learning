func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	bucket := make([][]int, len(nums)+1)
	for num, count := range counter {
		bucket[count] = append(bucket[count], num)
	}
	
	result := make([]int,0, k)
	for freq := len(bucket) - 1; freq >= 0 && len(result) < k; freq-- {
		for _, num := range bucket[freq] {
          result = append(result, num)
		  if len(result) == k {
			break
		  }
		}
	}

	return result
}
