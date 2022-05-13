package bits

// https://leetcode.com/problems/hamming-distance/

func hammingDistance(x int, y int) int {
	// xor operation 1&0=1, 0&0/1&1=0
	x = x ^ y

	count := 0
	for x > 0 {
		count++
		x &= x - 1
	}
	return count
}

func hammingDistance_1better(x int, y int) int {
	distance := 0

	for x > 0 || y > 0 {

		if (x&1)^(y&1) == 1 {
			distance++
		}
		x = x >> 1
		y = y >> 1
	}

	return distance
}

func hammingDistance_1(x int, y int) int {
	dist := 0

	for x > 0 && y > 0 {
		if x&1 != y&1 {
			dist++
		}
		x = x >> 1
		y = y >> 1
	}

	for x > 0 {
		if x&1 == 1 {
			dist++
		}
		x = x >> 1
	}

	for y > 0 {
		if y&1 == 1 {
			dist++
		}
		y = y >> 1
	}

	return dist
}
