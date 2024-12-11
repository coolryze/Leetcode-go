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
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}
