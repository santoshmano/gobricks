package dp

// Recursive solution
//
//         | 0, n==0
// fib(n) =| 1, n==1
//         | fib(n-1) + fib(n-2), n>1
//
// T(n) ~ O(2^n) --> number of nodes in recursion tree
// S(n) = O(n) 	 --> depth of recursion tree
//
func fib_rec(n int) int {
	if (n == 0) || (n == 1) {
		return n
	} else {
		return fib_rec(n-1) + fib_rec(n-2)
	}
}

// Memoized, Top down
// T(n) = O(n)
// S(n) = O(n)

func fib_memo(n int) int {

	if n < 2 {
		return n
	}
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	memo[0] = 0
	memo[1] = 1

	var _fib_memo func(n int) int

	_fib_memo = func(n int) int {
		if memo[n] == -1 {
			memo[n] = _fib_memo(n-1) + _fib_memo(n-2)
		}
		return memo[n]
	}

	return _fib_memo(n)
}

// dp, bottom up
// T(n) = O(n)
// T(n) = O(n)

func fib(n int) int {
	memo := make([]int, n+1)
	memo[0] = 0
	memo[1] = 1

	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}

// dp, bottom up, space optimized
// T(n) = o(n)
// S(n) = o(1)

func fib_opt(n int) int {
	memo := make([]int, 3)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i%3] = memo[(i-1)%3] + memo[(i-2)%3]
	}

	return memo[n%3]
}
