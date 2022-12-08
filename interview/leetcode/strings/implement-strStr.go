package strings

// https://leetcode.com/problems/implement-strstr/

func strStr(haystack string, needle string) int {

	if haystack == needle {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		j := 0
		for ; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}

	return -1
}
