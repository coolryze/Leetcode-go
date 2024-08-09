package targetsum

import "math"

func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	if int(math.Abs(float64(target))) > total {
		return 0
	}
	if (total+target)%2 == 1 {
		return 0
	}

	bagWeight := (total + target) / 2
	dp := make([]int, bagWeight+1)
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := bagWeight; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[bagWeight]
}
