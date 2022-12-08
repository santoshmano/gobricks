package math

import "math"

// prime number - a digit that has 2 positive divisors (usually 1 & itself)
// 0 - is not a prime as it has infinite divisors
// 1 - is not a prime as it only 1 positive divisor
//
// i/p - integer n
// o/p - number of primes from [0,n)

// approach 1
//	- for every number n,
//	-- check if it is a prime number by see if it is divisible
//	-- by any number from [2,sqrt(n)+1]
//
// t(n) - O(n*n)
func countPrimes(n int) int {

	isPrime := func(num int) bool {
		for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

// 12 - 2, 13
// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19
//
// approach 2
// - create an array/map for numbers from [2,n)
// - for every number i from [2,sqrt(n)+1]
// -- mark the multiples of i, from [i*i,i*i*i,....n) in the array
// - the unmarked numbers are primes.

func countPrimes(n int) int {

	count := 0
	notPrimeArr := make([]bool, n)

	for i := 2; i*i <= n; i++ {
		if notPrimeArr[i] == true {
			continue
		}

		for next := i + i; next < n; next += i {
			notPrimeArr[next] = true
		}
	}

	for i := 2; i < len(notPrimeArr); i++ {
		if notPrimeArr[i] == false {
			count++
		}
	}

	return count
}
