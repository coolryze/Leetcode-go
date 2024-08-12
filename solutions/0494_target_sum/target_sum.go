package targetsum

func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	diff := total - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	reg := diff / 2

	dp := make([]int, reg+1)
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := reg; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[reg]
}

func findTargetSumWays2(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	diff := total - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	reg := diff / 2

	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, reg+1)
	}
	dp[0][0] = 1

	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= reg; j++ {
			num := nums[i-1]
			if j >= num {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-num]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)][reg]
}

func findTargetSumWays3(nums []int, target int) int {
	res := 0

	var backtrack func(index, sum int)
	backtrack = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				res++
			}
		} else {
			backtrack(index+1, sum+nums[index])
			backtrack(index+1, sum-nums[index])
		}
	}

	backtrack(0, 0)

	return res
}
