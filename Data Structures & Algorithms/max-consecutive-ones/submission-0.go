func findMaxConsecutiveOnes(nums []int) int {
    var currRes, lastRes int


    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            currRes++
        } else {
            if currRes > lastRes {
            lastRes = currRes
            }
            currRes = 0
        }
    }
    if currRes > lastRes {
        return currRes
    }
    return lastRes
    
}