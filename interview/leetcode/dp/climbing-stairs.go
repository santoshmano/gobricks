package dp

// https://leetcode.com/problems/climbing-stairs

// Start from step 0 and go up to step N, climbing 1 step at a time
// or 2 steps at a time.
//
// i/p : n, steps to climb
// o/p : int, number of ways to climb upto step n
//
// impl approach 1- climb from step 0 to step n, or
// impl approach 2 - climb from step n to step 0
//
// recurrence relation = approach 1
//
//            | 0, k >= n
// cs(n, k) = | 1, k == n-1
//            | 2, k == n-2
//            | cs(n, k+1) + cs(n, k+2), k < n-2
// where, n == number of steps to climb
//        k == current step
//
// recurrence relation = approach 2
//
//         | 0, n == 0
// cs(n) = | 1, n == 1
//         | 2, n == 2
//         | cs(n-1) + cs(n-2), n > 2
// where, n == number of steps to climb DOWN
//        k == current step
//
// T(n) =
// S(n) =
//
func climbStairs_rec(n int) int {
	if n <= 2 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs_rec2(n int) int {
	var _cs func(k, n int) int

	_cs = func(n, k int) int {
		if k <= n-2 {
			return n - k // test
		}

		return _cs(n, k+1) + _cs(n, k+2)
	}

	return _cs(n, 0)
}

func climbStairsMemo(n int) int {
	if n <= 2 {
		return n
	}
	ways := make(map[int]int)
	ways[0] = 0
	ways[1] = 1
	ways[2] = 2

	var _cs func(int) int

	_cs = func(n int) int {

		if val, ok := ways[n]; ok {
			return val
		}

		if _, ok := ways[n-1]; !ok {
			ways[n-1] = _cs(n - 1)
		}

		if _, ok := ways[n-2]; !ok {
			ways[n-2] = _cs(n - 2)
		}

		return ways[n-1] + ways[n-2]
	}

	ways[n] = _cs(n)
	return ways[n]
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	ways := make([]int, n+1)
	ways[0] = 0
	ways[1] = 1
	ways[2] = 2

	for i := 3; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n]
}
