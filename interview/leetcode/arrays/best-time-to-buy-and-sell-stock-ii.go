package arrays

/*

- "you can buy it then immediately sell it on the same day."

examples :

`    7,1,5,3,6,4
    0 1 2 3 4 5
      ^   ^   ^
      `
    1 4

    - sell at immediate highest point
        repeat above by stepping to the right
algo :

repeat till idx < len(str)
    traverse till you find the lowest point, buy
    traverse till you find the highest point, sell
    accumulate profit
*/
func maxProfit(prices []int) int {

	var profit int = 0

	for idx := 0; idx < len(prices)-1; {

		// buy, find the lowest price
		for (idx < len(prices)-1) && prices[idx] >= prices[idx+1] {
			idx++
		}
		buy := idx

		// find the highest price
		for (idx < len(prices)-1) && prices[idx] <= prices[idx+1] {
			idx++
		}
		sell := idx

		profit += prices[sell] - prices[buy]
	}

	/* Test Case walkthrough
	   0 1 2 3 4 5
	   7,1,5,3,6,4
	     b s        p = 4, idx = 2
	         b s   p = 4+3, idx = 4
	*/

	return profit
}
