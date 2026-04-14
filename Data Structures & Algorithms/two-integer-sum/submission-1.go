func twoSum(nums []int, target int) []int {
    seen := map[int]int{}

	for idx, val := range nums {
		if _, ok := seen[val]; !ok {
			seen[val] = idx
		}
	}

	result := []int{}
	for idx, val := range nums {
		rem := target - val
		if res, ok := seen[rem]; ok {
			if idx == res {
				continue
			}
			result = append(result, res, idx)
			sort.Ints(result)
			break
		}
	}
	return result
}
