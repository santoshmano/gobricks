package word

// https://leetcode.com/problems/valid-palindrome/
func isAlphaNum(s uint8) bool {
	if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9') {
		return true
	} else {
		return false
	}
}

func toLower(lr uint8) uint8 {

	if lr >= 'A' && lr <= 'Z' {
		return lr + 32
	}
	return lr
}

func isPalindrome(s string) bool {

	for l, r := 0, len(s)-1; l < r; {
		if isAlphaNum(s[l]) {
			if isAlphaNum(s[r]) {
				if toLower(s[l]) == toLower(s[r]) {
					l = l + 1
					r = r - 1
				} else {
					//fmt.Println(string(toLower(s[l])), string(s[r]))
					return false
				}
			} else {
				r = r - 1
			}
		} else {
			l = l + 1
		}
	}

	return true
}
