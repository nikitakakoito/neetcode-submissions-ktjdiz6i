func singleNumber(nums []int) []int {
    xor := 0
    for _, num := range nums {
        xor ^= num
    }

    diff := xor & -xor

    a, b := 0, 0
    for _, num := range nums {
        if num&diff == 0 {
            a ^= num
        } else {
            b ^= num
        }
    }

    return []int{a, b}
}