package arrays

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {

	if len(prices) < 2 {
		return 0
	}

	minSofar := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {

		if (prices[i] - minSofar) > maxProfit {
			maxProfit = prices[i] - minSofar
		}

		if prices[i] < minSofar {
			minSofar = prices[i]
		}
	}

	return maxProfit
}
