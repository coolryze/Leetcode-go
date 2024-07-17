package jumpgameii

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	res := 0
	curDistance := 0
	nextDistance := 0

	for i := 0; i < len(nums)-1; i++ {
		nextDistance = max(nums[i]+i, nextDistance)
		if i == curDistance {
			res++
			curDistance = nextDistance
			if curDistance >= len(nums)-1 {
				break
			}
		}
	}

	return res
}
