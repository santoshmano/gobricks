package searching

// https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size
/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 *
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {

	// double your index, check if it exists
	// then if it is < or > target

	prevIndex := 0

	if reader.get(prevIndex) == target {
		return prevIndex
	}

	index := 1
	EXIT_VAL := 2147483647 // not used because it is same as target

	for prevIndex < index {

		val := reader.get(index)

		if target > val {
			prevIndex = index
			index = index * 2
		} else if (target < val) || (val == EXIT_VAL) {
			index = prevIndex + (index-prevIndex)/2
			if index <= prevIndex {
				return -1
			}
		} else {
			// val == target
			return index
		}
	}

	return -1
}
