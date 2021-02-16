package sorting

/*
Insert an element into a sorted array, repeat.
hint - array of length 0/1, is already sorted

  - given a set of elements ranging (i, n), pick the ith element
    and insert it into the sorted range (0,i-1), by moving the bigger
    elements one step to the right.

  - repeat the above for range (i+1, n)

  T(n) - O(n**2)
  S(n) - O(1)
  Stable sort - Yes
*/
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {

		// find correct pos for ith elem in range (0,i-1)
		numToInsert := nums[i]

		// shift all bigger elements one step right
		j := i - 1
		for ; (j >= 0) && (numToInsert < nums[j]); j-- {
			nums[j+1] = nums[j]
		}

		// if any shift occurred then insert
		if j < i-1 {
			nums[j+1] = numToInsert
		}
	}
}

// Recursive insertion sort
func InsertionSortRec(nums []int) {
	insertionSortRec(nums, 0, len(nums)-1)
}

func insertionSortRec(nums []int, start int, end int) {
	if start > end {
		return
	}

	// array from index 0 to start-1 is sorted, place nums[start]
	// in that range
	val := nums[start]
	i := start - 1
	for ; i >= 0 && nums[i] > val; i-- {
		nums[i+1] = nums[i]
	}
	nums[i+1] = val

	insertionSortRec(nums, start+1, end)
}
