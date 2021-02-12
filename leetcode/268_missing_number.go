package leetcode

func missingNumber(nums []int) int {
	totalSum := len(nums) * (len(nums) + 1) / 2
	for _, num := range nums {
		totalSum -= num
	}
	return totalSum

}
