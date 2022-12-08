package arrays

func rotate1(nums []int, k int) {

	/*
		Time Complexity - O(n)
		 Space Complexity - O(k)
	*/

	// for cases when k > len(nums)
	k = k % len(nums)

	// Create an extra array of size k
	arr := make([]int, k)

	// & copy the last k items into it

	i, j := 0, len(nums)-k
	for (i < k) && (j < len(nums)) {
		arr[i] = nums[j]
		i++
		j++
	}

	// Move the first n-k items k steps to the right
	i, j = len(nums)-k-1, len(nums)-1
	for i >= 0 {
		nums[j] = nums[i]
		j--
		i--
	}

	// Copy back the first set of k items into the position 0 to k-1
	i, j = 0, 0
	for j < k {
		nums[i] = arr[j]
		i++
		j++
	}

}

/*

 0 1 2 3 4 5 6, k=3
 1 2 3 4 5 6 7
       1
             4
     7     3

*/

func rotate(nums []int, k int) {

	// move n items, to a new position
	//  - (curpos+k)%n

	i := 0
	k = k % len(nums)
	if k == 0 {
		return
	}
	curPos := i
	curVal := nums[i]

	count := len(nums)

	for i < len(nums) && count > 0 {

		nextPos := (curPos + k) % len(nums)
		nextVal := nums[nextPos]

		nums[nextPos] = curVal

		if nextPos == i {
			i++
			curPos = i
			curVal = nums[i]
		} else {
			curPos = nextPos
			curVal = nextVal
		}
		count--
	}

}
