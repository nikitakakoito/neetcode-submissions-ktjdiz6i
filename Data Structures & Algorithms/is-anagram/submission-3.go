func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}


	seenS := map[rune]int{}
	seenT := map[rune]int{}

	for _, vals := range s {
		seenS[vals]++
	}

	for _, valt := range t {
		seenT[valt]++
	}

	for key, val := range seenS {
		if v, ok := seenT[key]; !ok {
			return false
		} else if v != val {
			return false
		}
	}
	return true
}