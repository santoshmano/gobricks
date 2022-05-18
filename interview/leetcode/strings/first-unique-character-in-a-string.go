package strings

func firstUniqChar(s string) int {

	for i := 0; i < len(s); i++ {

		found := false
		for j := 0; j < len(s); j++ {
			if (j != i) && (s[j] == s[i]) {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}

	return -1
}

/*	// map version
func firstUniqChar(s string) int {

	smap := make(map[rune]int)

	for _, val := range s {
		smap[val] += 1
	}

	for i, val := range s {
		if smap[val] == 1 {
			return i
		}
	}
	return -1
}
*/
