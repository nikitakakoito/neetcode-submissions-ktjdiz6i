func singleNumber(nums []int) int {
	хуй := 0
	for _, val := range nums {
		хуй ^= val
	}
	return хуй
}
