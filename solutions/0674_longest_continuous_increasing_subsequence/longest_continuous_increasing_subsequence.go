package longestcontinuousincreasingsubsequence

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = 1
	}
	res := dp[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
