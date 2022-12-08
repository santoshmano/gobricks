package bits

// https://leetcode.com/problems/number-of-1-bits/

func hammingWeight2(num uint32) int {
	// can generate it instead of assigning
	weightMap := map[uint32]int{
		uint32(0b0001): 1,
		uint32(0b0010): 1,
		uint32(0b0011): 2,
		uint32(0b0100): 1,
		uint32(0b0101): 2,
		uint32(0b0110): 2,
		uint32(0b0111): 3,
		uint32(0b1000): 1,
		uint32(0b1001): 2,
		uint32(0b1010): 2,
		uint32(0b1100): 2,
		uint32(0b1011): 3,
		uint32(0b1101): 3,
		uint32(0b1110): 3,
		uint32(0b1111): 4,
	}

	weight := 0

	for num > 0 {
		weight += weightMap[num&0xF]
		num >>= 0xF
	}

	return weight
}

func hammingWeight_1(num uint32) int {
	oneCount := 0
	for num > 0 {
		if (num & 1) == 1 {
			oneCount++
		}
		num = num >> 1
	}
	return oneCount
}

func hammingWeight(num uint32) int {
	count := 0

	for num > 0 {
		count++
		num = num & (num - 1) //removes last set bit
	}
	return count
}
