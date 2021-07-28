package stacks

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {

	var stk []rune

	for _, val := range s {
		switch val {
		case ')':
			if len(stk) == 0 {
				return false
			}
			if stk[len(stk)-1] != '(' {
				return false
			}
			stk = stk[:len(stk)-1]
		case ']':
			if len(stk) == 0 {
				return false
			}
			if stk[len(stk)-1] != '[' {
				return false
			}
			stk = stk[:len(stk)-1]
		case '}':
			if len(stk) == 0 {
				return false
			}
			if stk[len(stk)-1] != '{' {
				return false
			}
			stk = stk[:len(stk)-1]
		default:
			stk = append(stk, val)
		}
	}

	if len(stk) != 0 {
		return false
	}

	return true

}
