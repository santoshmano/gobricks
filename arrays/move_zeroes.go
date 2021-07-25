package arrays

//https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	readIndex := 0
	writeIndex := 0

	for readIndex < len(nums) {
		if nums[readIndex] == 0 {
			readIndex++
		} else {
			nums[writeIndex] = nums[readIndex]
			readIndex++
			writeIndex++
		}
	}

	for writeIndex < len(nums) {
		nums[writeIndex] = 0
		writeIndex++
	}
}
