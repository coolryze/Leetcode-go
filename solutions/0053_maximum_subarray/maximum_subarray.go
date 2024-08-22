package maximumsubarray

import (
	"math"
)

func maxSubArray(nums []int) int {
	res := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}

func maxSubArray2(nums []int) int {
	res := math.MinInt

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > res {
			res = sum
		}
		if sum <= 0 {
			sum = 0
		}
	}

	return res
}
