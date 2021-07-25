package bits

// https://leetcode.com/problems/number-of-1-bits/

func hammingWeight(num uint32) int {
	oneCount := 0
	for num > 0 {
		if (num & 1) == 1 {
			oneCount++
		}
		num = num >> 1
	}
	return oneCount
}
