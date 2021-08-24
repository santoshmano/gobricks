package recursion

// https://leetcode.com/problems/letter-case-permutation/

import "strings"

func letterCasePermutation(S string) []string {

	var result []string

	var _letterCP func(idx int, tmp string)

	_letterCP = func(idx int, tmp string) {
		if idx >= len(S) {
			result = append(result, tmp)
			return
		}

		letter := S[idx]
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			_letterCP(idx+1, tmp+strings.ToLower(string(letter)))
			_letterCP(idx+1, tmp+strings.ToUpper(string(letter)))
		} else {
			_letterCP(idx+1, tmp+string(letter))
		}

	}

	_letterCP(0, "")
	return result
}

// bytes library?

func letterCasePermutationS(S string) []string {

	var result []string

	b := []byte(S)

	//var tmp []byte
	tmp := make([]byte, 0, len(S))
	var _letterCP func(idx int)

	_letterCP = func(idx int) {
		if idx >= len(S) {
			result = append(result, string(tmp))
			return
		}

		letter := b[idx]
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			tmp = append(tmp, strings.ToLower(string(letter))[0])
			_letterCP(idx + 1)
			tmp = tmp[:len(tmp)-1]
			tmp = append(tmp, strings.ToUpper(string(letter))[0])
			_letterCP(idx + 1)
			tmp = tmp[:len(tmp)-1]
		} else {
			tmp = append(tmp, letter)
			_letterCP(idx + 1)
			tmp = (tmp)[:len(tmp)-1]
		}
	}

	_letterCP(0)
	return result
}

func letterCasePermutationS(S string) []string {

	var result []string

	b := []byte(S)

	var _letterCP func(idx int, tmp *[]byte)

	_letterCP = func(idx int, tmp *[]byte) {
		if idx >= len(S) {
			result = append(result, string(*tmp))
			return
		}

		letter := b[idx]
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			*tmp = append(*tmp, strings.ToLower(string(letter))[0])
			_letterCP(idx+1, tmp)
			*tmp = (*tmp)[:len(*tmp)-1]
			*tmp = append(*tmp, strings.ToUpper(string(letter))[0])
			_letterCP(idx+1, tmp)
			*tmp = (*tmp)[:len(*tmp)-1]
		} else {
			*tmp = append(*tmp, letter)
			_letterCP(idx+1, tmp)
			*tmp = (*tmp)[:len(*tmp)-1]
		}
	}

	_letterCP(0, &[]byte{})
	return result
}
