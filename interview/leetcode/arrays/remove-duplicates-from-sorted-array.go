package arrays

/*

examples 1
  1 2 3 4
0 0 1 1 1 2 2 3 3 4
                  r
        w

examples 2
        5
0 1 2 3 3 5 6
            r
        w


if nums[r] != nums[r-1] {
    if r = w+1 {
        r+
        w++
    } else {

    }
}
*/

func removeDuplicates(nums []int) int {
	read, write := 1, 0

	for read < len(nums) {
		if nums[read] != nums[read-1] {
			if read == write+1 {
				read++
				write++
			} else {
				nums[write+1] = nums[read]
				write++
				read++
			}
		} else {
			read++
		}
	}

	return write + 1
}

// old impl
func removeDuplicates1(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	// have a read & write index
	read := 1
	write := 1
	for read < len(nums) {
		if nums[read] != nums[read-1] {
			nums[write] = nums[read]
			write++
		}
		read++
	}

	return write
}
