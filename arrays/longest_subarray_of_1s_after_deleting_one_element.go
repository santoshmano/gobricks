package arrays

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/submissions/

// i/p - array of numbers, with  value 1 or 0
// o/p - int , length of longest non-0 subarray if we remove one 0

// approach 1
// traverse array till end
//      find length of zeroes
//      find length of cur 1's
//      if length of 0s was 1
//          maxSub = max(maxSum, prev+cur)
//      else
//          maxSum = max(maxSum, cur)
//      reset pre, cur, zero values

func longestSubarray(nums []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	// ex1 - [ 0 0 0 0]
	//  0 1 2 3 4 5 6 7 8
	// ex2 - [0,1,1,1,0,1,1,0,1]
	// idx = 1, sub1 = 3, idx = 4 -- idx = 5, sub2 = 2, idx = 7, maxSub = 6
	//

	zero, prev, cur, maxSub := 0, 0, 0, 0
	anyZero := false

	for idx := 0; idx < len(nums); {

		zero = 0
		cur = 0
		// get the zero length
		for idx < len(nums) && nums[idx] == 0 {
			anyZero = true
			zero = zero + 1
			idx++
		}

		// get the first subarray length
		for idx < len(nums) && nums[idx] == 1 {
			cur++
			idx++
		}

		if zero < 2 {
			maxSub = max(maxSub, cur+prev)
		} else {
			maxSub = max(maxSub, cur)
		}
		prev = cur
	}

	// for condition if all are '1's, then got to delete one item
	if maxSub > 0 && !anyZero {

		return maxSub - 1
	} else {
		return maxSub

	}
}
