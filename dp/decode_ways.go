package dp

import "strconv"

// https://leetcode.com/problems/decode-ways/

// Given an encoded string, return the number of possible valid decodings.

// i/p = string of numbers
// o/p = int, number of possible valid decodings

// branching factor of tree && height of tree
// T(n) -
// S(n) -
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

		// if choosing one character
		ways := _numDecodings(idx + 1)

		// if choosing two characters
		if idx < (len(s) - 1) {
			num, _ := strconv.Atoi(s[idx : idx+2]) // example s[1:3] == "26"
			if num >= 10 && num <= 26 {
				ways += _numDecodings(idx + 2)
			}
		}

		return ways
	}

	return _numDecodings(0)
}

// memo

// Given an encoded string, return the number of possible valid decodings.

// i/p = string of numbers
// o/p = int, number of possible valid decodings

func numDecodingsMemo(s string) int {

	if len(s) == 0 {
		return 0
	}

	// key = idx, val = ways
	memo := make(map[int]int)

	var _numDecodings func(int) int

	_numDecodings = func(idx int) int {

		if idx >= len(s) {
			return 1
		}

		// if the first string is a 0, then it is invalid
		if s[idx] == '0' {
			return 0
		}

		if val, ok := memo[idx]; ok {
			return val
		}

		// if choosing one character
		memo[idx] = _numDecodings(idx + 1)

		// if choosing two characters
		if idx < (len(s) - 1) {
			num, _ := strconv.Atoi(s[idx : idx+2]) // example s[1:3] == "26"
			if num >= 10 && num <= 26 {
				memo[idx] += _numDecodings(idx + 2)
			}
		}

		return memo[idx]
	}

	// speed of coding && time complexity

	return _numDecodings(0)
}
