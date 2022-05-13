package dp

/*

        0. 1.  2  3.  4. 5. 6. 7.  8
ex1   [-2  1  -3  4  -1  2  1  -5  4]

left    right   sum maxsum
0       0       0   0
0       1       -2  0
1       1       0   0
1       2       -3  0
2       2       0   0
2       3       4   4
2       6       6   6
2       7       1   6
2       8       5   6

ex2  [5, 4, -1, 7, 8]
      0, 1,  2, 3, 4

left    right   sum maxsum
0       0       0   0
0       2       9   9
0       3       8   9
0       5       23  23

*/

// # approach 1
//
// calculate sum of all subarrays
// return highest
//
// T(n) - O(n^2)
// S(n) - O(1)
//
// # approach 2
//
// 	- use two pointers, left & right
// 	- increment right pointer and maintain running sum and max sum, as long as it is +ve
// 	- if sum is -ve then increment left pointer till it is +ve and repeat
// 	--	repeat above step till right reaches the end.
//
// T(n) = O(n)
// S(n) = O(1)
//
func maxSubArray(nums []int) int {

	left, right := 0, 0
	maxSum := -1 << 31
	sum := 0

	for left < len(nums) && right < len(nums) {

		sum += nums[right]

		if sum > maxSum {
			maxSum = sum
		}

		right++

		if sum <= 0 {
			left = right
			sum = 0
		}

	}

	return maxSum
}
