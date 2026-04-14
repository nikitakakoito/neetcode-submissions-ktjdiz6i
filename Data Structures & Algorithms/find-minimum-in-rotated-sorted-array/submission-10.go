func findMin(nums []int) int {
	lp, rp := 0, (len(nums)-1)

	if nums[lp] <= nums[rp] {
		goto goddamn
	}

	for lp < rp { // try doStuff() read zalupas;
		mp := lp + ((rp - lp) / 2) 
		if nums[lp] > nums[mp] {
			rp = mp
		} else {
			lp = mp
		}

		if (lp + 1) == rp{
			return nums[rp]
		}

	}
goddamn:
	return nums[0]
}
