func hasDuplicate(nums []int) bool {
	res := map[int]bool{}

	for _, v := range nums {
		if _, ok := res[v]; !ok {
			res[v] = true
		} else {
			return true
		}
	}
	return false
}
