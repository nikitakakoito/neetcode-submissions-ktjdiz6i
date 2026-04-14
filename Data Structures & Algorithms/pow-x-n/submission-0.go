func myPow(x float64, n int) float64 {
	// var res = x
	// for i := 1; i < n; i++ {
	// 	res *= x
	// }
	// return res
	return math.Pow(x, float64(n))
}