package partitionequalsubsetsum

func canPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}
	target := total / 2

	dp := make([]int, target+1)

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[target] == target
}

func canPartition2(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	var backtracking func(nums []int, target int, index int) bool
	backtracking = func(nums []int, target int, index int) bool {
		if target == 0 {
			return true
		}
		if target < 0 || index == len(nums) {
			return false
		}

		if backtracking(nums, target-nums[index], index+1) {
			return true
		}

		return backtracking(nums, target, index+1)
	}

	return backtracking(nums, total/2, 0)
}
