package word

// https://leetcode.com/problems/reverse-integer/
import "math"

func reverse(x int) int {
	minus := false
	if x < 0 {
		minus = true
		x = x * -1
	}

	y := 0
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}

	if y > int(math.Pow(2, 31)-1) {
		return 0
	}

	if minus {
		return y * -1
	} else {
		return y
	}
}
