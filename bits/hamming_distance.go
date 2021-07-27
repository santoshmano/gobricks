package bits

// https://leetcode.com/problems/hamming-distance/

func hammingDistance(x int, y int) int {

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
