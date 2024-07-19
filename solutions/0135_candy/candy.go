package candy

func candy(ratings []int) int {
	nums := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		nums[i] = 1
	}

	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			nums[i+1] = nums[i] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			nums[i-1] = max(nums[i]+1, nums[i-1])
		}
	}

	res := 0
	for i := 0; i < len(ratings); i++ {
		res += nums[i]
	}
	return res
}
