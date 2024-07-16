package jumpgame

func canJump(nums []int) bool {
	cover := 0

	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)
		if cover >= len(nums)-1 {
			return true
		}
	}

	return false
}
