package arrays

// https://leetcode.com/problems/decrease-elements-to-make-array-zigzag/

func movesToMakeZigzag(nums []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	// example 1
	// len(nums) = 5
	//  0. 1. 2. 3. 4
	// [9, 6, 1, 6, 2]

	// example 2
	// len(nums) = 4
	//  0. 1. 2. 3
	// [1, 3, 4, 6]

	// check if odd-indexed is lower, if not note how much
	// to decrement
	oddCount := 0
	idx := 1
	for idx < len(nums) {
		left, right := 0, 0
		if nums[idx] >= nums[idx-1] {
			left = (nums[idx] - nums[idx-1]) + 1
		}

		if (idx+1) < len(nums) && nums[idx] >= nums[idx+1] {
			right = (nums[idx] - nums[idx+1]) + 1
		}

		oddCount += max(left, right)
		idx += 2
	}

	// check if even-indexed is lower, if not note how much
	// to decrement
	evenCount := 0
	idx = 0
	for idx < len(nums) {
		left, right := 0, 0
		if ((idx - 1) >= 0) && (nums[idx] >= nums[idx-1]) {
			left = (nums[idx] - nums[idx-1]) + 1
		}

		if ((idx + 1) <= len(nums)-1) && nums[idx] >= nums[idx+1] {
			right = (nums[idx] - nums[idx+1]) + 1
		}

		evenCount += max(left, right)
		idx += 2
	}

	return min(oddCount, evenCount)
}
