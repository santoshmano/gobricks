package arrays

func generate(numRows int) [][]int {
	var result [][]int
	if numRows == 0 {
		return [][]int{}
	}

	result = append(result, []int{1})
	for i := 1; i < numRows; i++ {
		prevRow := result[len(result)-1]
		curRow := []int{1}

		for i := 1; i < len(prevRow); i++ {
			curRow = append(curRow, prevRow[i]+prevRow[i-1])
		}
		curRow = append(curRow, 1)
		result = append(result, curRow)
	}

	return result
}
