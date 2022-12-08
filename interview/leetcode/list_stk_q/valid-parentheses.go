package list_stk_q

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {

	// use a slice as a stack
	// top - stk[len(s)-1]
	var stk []byte
	match := map[byte]byte{')': '(', '}': '{', ']': '['}

	for i := 0; i < len(s); i++ {

		switch s[i] {

		case '(', '{', '[':
			stk = append(stk, s[i])

		case ')', '}', ']':
			if len(stk) == 0 {
				return false
			}

			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]

			if top != match[s[i]] {
				return false
			}

		default:
			// unexpected character
			return false
		}
	}

	if len(stk) == 0 {
		return true
	}

	return false
}
