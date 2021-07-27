package arrays

// https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	singleNum := nums[0]
	for i := 1; i < len(nums); i++ {
		singleNum ^= nums[i]
	}

	return singleNum
}
