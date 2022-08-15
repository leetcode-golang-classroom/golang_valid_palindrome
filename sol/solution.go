package sol

func isPalindrome(s string) bool {
	lp, rp := 0, len(s)-1
	var isAlphaBet = func(c byte) bool {
		if ('a' <= c && 'z' >= c) || ('A' <= c && 'Z' >= c) || ('0' <= c && c <= '9') {
			return true
		}
		return false
	}
	var toLower = func(c byte) byte {
		if 'A' <= c && 'Z' >= c {
			return (c - 'A') + 'a'
		}
		return c
	}
	for lp <= rp {
		if !isAlphaBet(s[lp]) {
			lp++
			continue
		}
		if !isAlphaBet(s[rp]) {
			rp--
			continue
		}
		if toLower(s[rp]) != toLower(s[lp]) {
			return false
		}
		rp--
		lp++
	}
	return true
}
