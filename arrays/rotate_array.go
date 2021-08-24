package arrays

func rotate(nums []int, k int) {

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
