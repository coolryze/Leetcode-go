package laststoneweightii

func lastStoneWeightII(stones []int) int {
	total := 0
	for _, weight := range stones {
		total += weight
	}

	target := total / 2
	dp := make([]int, target+1)

	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return total - dp[target]*2
}
