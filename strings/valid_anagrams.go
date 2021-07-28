package word

// https://leetcode.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {

	smap := make(map[rune]int)

	for _, srune := range s {
		smap[srune]++
	}

	for _, trune := range t {
		smap[trune]--
	}

	for _, value := range smap {
		if value != 0 {
			return false
		}
	}
	return true
}
