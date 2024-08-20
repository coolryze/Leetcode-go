package besttimetobuyandsellstockiv

func maxProfit(k int, prices []int) int {
	indexNum := k*2 + 1
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, indexNum)
	}
	for j := 1; j < indexNum; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j < indexNum; j++ {
			if j%2 == 1 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}

	return dp[len(prices)-1][indexNum-1]
}
