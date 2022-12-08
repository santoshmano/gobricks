package arrays

//https://leetcode.com/problems/move-zeroes/

func moveZeroes1(nums []int) {
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

func moveZeroes(nums []int) {
	read, write := 0, 0

	for read < len(nums) {
		if nums[read] != 0 {
			nums[write] = nums[read]
			write++
		}
		read++
	}

	for write < len(nums) {
		nums[write] = 0
		write++
	}
}
