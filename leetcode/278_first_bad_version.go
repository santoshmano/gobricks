package leetcode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {

	start := 1
	end := n

	for start <= end {
		mid := start + (end-start)/2

		if isBadVersion(mid) == true {

			if mid-1 >= 1 {
				if isBadVersion(mid-1) == true {
					end = mid - 1
				} else {
					return mid
				}
			} else {
				return mid
			}

		} else {
			start = mid + 1
		}
	}
	return -1
}

// assumption that there is always a
func firstBadVersionOptimized(n int) int {

	start := 1
	end := n

	for start <= end {
		mid := start + (end-start)/2

		if isBadVersion(mid) == true {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return start
}
