package arrays

func plusOne(digits []int) []int {

	sum := make([]int, len(digits)+1)
	remainder := 1 // as a plus one addition is needed

	for i := len(digits) - 1; i >= 0; i-- {

		val := digits[i] + remainder

		sum[i+1] = val % 10
		remainder = val / 10
	}

	if remainder > 0 {
		sum[0] = 1
	} else {
		sum = sum[1:]
	}

	return sum
}
