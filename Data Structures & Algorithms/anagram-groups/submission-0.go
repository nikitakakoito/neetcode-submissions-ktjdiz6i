func groupAnagrams(strs []string) [][]string {
	freq := make(map[[26]int][]string)
	for _, word := range strs {
		var mask [26]int
		for _, ch := range word {
			mask[ch-'a']++
		}
		freq[mask] = append(freq[mask], word)
	}
	
	res := make([][]string, 0)
	for _, v := range freq {
		res = append(res, v)
	}
	return res
}
