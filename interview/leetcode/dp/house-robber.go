package dp

import "fmt"

// https://leetcode.com/problems/house-robber/

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

// tabulation, bottom up
func rob(nums []int) int {

	maxLoot := make([]int, len(nums)+1)

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	maxLoot[0] = 0
	maxLoot[1] = nums[0]
	maxLoot[2] = max(nums[0], nums[1])

	for i := 3; i < len(nums)+1; i++ {
		maxLoot[i] = nums[i-1] + max(maxLoot[i-2], maxLoot[i-3])
	}

	fmt.Println(maxLoot)
	return max(maxLoot[len(nums)-1], maxLoot[len(nums)])
}

// memoization
func robMemo(nums []int) int {

	memo := make([]int, len(nums))

	for i, _ := range memo {
		memo[i] = -1
	}

	var _rob func(int) int

	_rob = func(i int) int {
		if i == len(nums)-1 {
			return nums[i]
		} else if i >= len(nums) {
			return 0
		}

		if memo[i] == -1 {
			memo[i] = nums[i] + max(_rob(i+2), _rob(i+3))
		}

		return memo[i]
	}

	return max(_rob(0), _rob(1))
}

// brute force, recursive solution
//
// rob(i) = 0, i >= len(nums)			// if there are no houses
// 		  = nums[i], i == len(nums)- 1	// if there is only 1 house
//        = nums[i] + max(rob(i+2), rob(i+3))
//					//choose the first house or the second house, and recurse
func robR(nums []int) int {

	var _rob func(int) int

	_rob = func(i int) int {

		if i == len(nums)-1 {
			return nums[i]
		} else if i >= len(nums) {
			return 0
		}

		return nums[i] + max(_rob(i+2), _rob(i+3))
	}

	return max(_rob(0), _rob(1))
}
