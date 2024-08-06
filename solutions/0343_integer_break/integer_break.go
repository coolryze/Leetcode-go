package integerbreak

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(max(j*(i-j), j*dp[i-j]), dp[i])
		}
	}

	return dp[n]
}
