package sorting

// https://leetcode.com/problems/missing-number/

// using cycle sort technique
func missingNumber(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			val := nums[i]
			if val == len(nums) {
				break
			}
			nums[i], nums[val] = nums[val], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return len(nums)

}

// using cycle sort technique
func missingNumber1(nums []int) int {

	for i := 0; i < len(nums); i++ {

		// work on it this index till we get the right value
		for j := i; nums[j] != j; {

			// if it is n, break
			if nums[j] == len(nums) {
				break
			}

			//place the value at correct index, and swap with tht index
			nums[j], nums[nums[j]] = nums[nums[j]], nums[j]
		}
	}

	// now array is sorted, if any index has value n, that index is missing its value
	for i, v := range nums {
		if v == len(nums) {
			return i
		}
	}

	// else the value n is misssing.
	return len(nums)
}

// math technique; not favourable
func missingNumber2(nums []int) int {

	n := len(nums)
	sum := 0

	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	return n*(n+1)/2 - sum
}
