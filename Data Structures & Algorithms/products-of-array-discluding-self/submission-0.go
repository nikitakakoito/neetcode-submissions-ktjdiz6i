func productExceptSelf(nums []int) []int {

	res := make([]int, len(nums))
	for i, _ := range nums {
		res[i] = 1
  		for ii, _ := range nums {
     		if i != ii {
     		res[i] *= nums[ii]
     		}
    	}
	}
	return res
}
