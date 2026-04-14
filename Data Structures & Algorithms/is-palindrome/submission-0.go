func isPalindrome(s string) bool {
	lp, rp := 0, len(s)-1

	for lp < rp {
		for lp < rp && !isAlphaNumeric(s[lp]) {
			lp++
		}
		for lp < rp && !isAlphaNumeric(s[rp]) {
			rp--
		}

		lc, rc := s[lp], s[rp]

		if lc >= 'A' && lc <= 'Z' {
			lc += 32
		}

		if rc >= 'A' && rc <= 'Z' {
			rc += 32
		}

		if lc != rc {
			return false
		}
		lp++
		rp--
	}
	return true
}


func isAlphaNumeric(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}