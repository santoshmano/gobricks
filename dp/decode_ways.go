package dp

import "strconv"

// Given an encoded string, return the number of possible valid decodings.

// i/p = string of numbers
// o/p = int, number of possible valid decodings

func numDecodings(s string) int {

	// boundary conditions
	// len(s) == 0
	// s has non integer characters

	// Example - "226" == 3
	// idx = 0, return _n(1)+_n(2)

	if len(s) == 0 {
		return 0
	}

	var _numDecodings func(int) int

	_numDecodings = func(idx int) int {

		if idx >= len(s) {
			return 1
		}

		// if the first string is a 0, then it is invalid
		if s[idx] == '0' {
			return 0
		}

		if idx < len(s)-1 {

			num, _ := strconv.Atoi(s[idx : idx+2]) // example s[1:3] == "26"
			if num >= 10 && num <= 26 {
				return _numDecodings(idx+1) + _numDecodings(idx+2)
			}
		}

		// make it more concise
		return _numDecodings(idx + 1)
	}

	// branching factor of tree && height of tree

	// speed of coding && time complexity

	return _numDecodings(0)
}
