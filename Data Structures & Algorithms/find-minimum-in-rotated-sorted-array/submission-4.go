func findMin(nums []int) int {
	lp, rp := 0, (len(nums)-1)

	if nums[lp] <= nums[rp] {
		return nums[lp]
	}

	for lp <= rp { // try doStuff() read zalupas;
		mp := lp + ((rp - lp) / 2) // middle pointer
		// 0+((2-0)/2)
		if nums[lp] > nums[mp] {
			rp = mp
		}
		if nums[mp] > nums[rp] {
			lp = mp
		}

		if ((lp + 1) == rp) && (mp == rp) {
			return nums[rp]
		}

		if ((rp - 1) == lp) && (mp == lp) {
			return nums[rp]
		}



		// if nums[lp] < nums[lp-1] {
		// 	return nums[lp]
		// }

		// if nums[lp] > nums[lp+1]{
		// 	return nums[lp+1]
		// }

		// if nums[mp] > nums[rp] {
		// 	return nums[rp]
		// }
	}

	return nums[0]
}
