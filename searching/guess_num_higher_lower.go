package searching

// https://leetcode.com/problems/guess-number-higher-or-lower/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {

	var (
		start int = 0
		end   int = n
		mid   int = 0
	)

	for start <= end {
		mid = start + (end-start)/2

		switch guess(mid) {
		case 0:
			return mid
		case 1:
			start = mid + 1
		case -1:
			end = mid - 1
		}
	}

	return -1
}

func main() {

}
