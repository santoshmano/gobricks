package dp

// i/p - prices array
// o/p - maximum profit by buy and selling one stock
//  [7,1,5,3,6,4]
//   0 1 2 3 4 5

// approach 1
// 	- buy one day 0, and find the best day to sell, store the maxProfit
// 	- repeat the same to buy on the next day and update maxProfit if needed
//
//	t(n) = O(n*n)
//  s(n) = O(1)
//
func maxProfit(prices []int) int {

	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}

	if len(prices) < 2 {
		return 0
	}

	// find the stock that is biggest
	maxP := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[0] > maxP {
			maxP = prices[i] - prices[0]
		}
	}

	return max(maxP, maxProfit(prices[1:]))
}

// approach 2
//	- find/store the stock that is largest after current day.
//  - traverse the array and keep track of maxprofit
//
// t(n) - O(n)
// s(n) - O(n)
func maxProfitMemo(prices []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}

	maxAfter := make([]int, len(prices))

	maxSoFar := 0
	for i := len(prices) - 1; i >= 0; i-- {
		maxSoFar = max(maxSoFar, prices[i])
		maxAfter[i] = maxSoFar
	}

	maxProf := 0

	for i := 0; i < len(prices); i++ {
		maxProf = max(maxProf, maxAfter[i]-prices[i])
	}

	return maxProf
}

// approach 3
// - keep note of min stock price so far
// - traverse the array and calculate maxProfit as we know minstock price so far
//
// t(n) = O(n)
// s(n) = O(1)
func maxProfit(prices []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	min := func(i, j int) int {
		if i < j {
			return i
		} else {
			return j
		}
	}
	minSofar := prices[0]
	maxProf := 0
	for i := 1; i < len(prices); i++ {
		minSofar = min(minSofar, prices[i])
		maxProf = max(maxProf, prices[i]-minSofar)
	}

	return maxProf
}
