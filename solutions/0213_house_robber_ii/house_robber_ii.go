package houserobberii

func rob(nums []int) int {
	if len(nums) <= 1 {
		return robWithoutCircle(nums)
	}

	res1 := robWithoutCircle(nums[:len(nums)-1])
	res2 := robWithoutCircle(nums[1:])

	return max(res1, res2)
}

func robWithoutCircle(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]

	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}
