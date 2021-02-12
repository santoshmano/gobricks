package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {

	start := 0
	end := len(letters) - 1

	for start <= end {
		mid := start + (end-start)/2

		if letters[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}

	return letters[start%len(letters)]
}
