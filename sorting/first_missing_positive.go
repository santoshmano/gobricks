package sorting

// https://leetcode.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {

	if len(nums) == 0 {
		return 1
	}

	// sort it using cycle sort
	for i := 0; i < len(nums); i++ {

		for nums[i] != i+1 {

			if (nums[i] < 1) || (nums[i] > len(nums)-1) || ((nums[i] > 0) && (nums[i] == nums[nums[i]-1])) {
				break
			} else {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}

		}

	}

	// find the smallest number
	minval := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == minval {
			minval = nums[i] + 1
		} else {
			break
		}

	}
	return minval

}
