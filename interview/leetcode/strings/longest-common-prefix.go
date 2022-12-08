package strings

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	count := 0
	for count < len(strs[0]) {
		for idx := 1; idx < len(strs); idx++ {
			if (count >= len(strs[idx])) || strs[idx][count] != strs[0][count] {
				return strs[0][:count]
			}
		}
		count++
	}
	return strs[0][:count]
}
