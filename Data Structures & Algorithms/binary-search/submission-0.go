func search(nums []int, target int) int {
	lp, rp := 0, len(nums)-1
	for lp <= rp {
		mp := lp + (rp-lp)/2
		if nums[mp] == target {
			return mp
		}
		
		if nums[mp] < target {
			lp = mp+1
		} else {
			rp = mp-1
		}
	}
	return -1
}
