package dp
// https://leetcode.com/problems/min-cost-climbing-stairs/
xx
// minCost(n) = min(minCost(n-1), minCost(n-2)) + cost(n-1)
// T(n) = O(2^n)
// S(n) = O(n)

// it has an optimal substructure, where the prefix cost is also optimal

// ex
// idx         0   1.  2.  3
// cost    = [10, 15, 20]
// minCost = [10, 15, 30, 15 ]
// len(cost)    = 3
// len(minCost) = 4

// Bottomup

func minCostClimbingStairs(cost []int) int {

	minCost := make([]int, len(cost)+1)

	minCost[0] = cost[0]
	minCost[1] = cost[1]

	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}

	for i := 2; i < len(minCost)-1; i++ {
		minCost[i] = min(minCost[i-1], minCost[i-2]) + cost[i]
	}
	//fmt.Println(minCost)
	return min(minCost[len(minCost)-2], minCost[len(minCost)-3])
}
