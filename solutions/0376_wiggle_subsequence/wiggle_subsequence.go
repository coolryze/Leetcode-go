package wigglesubsequence

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	res := 1
	curDiff := 0
	prevDiff := 0
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if (curDiff > 0 && prevDiff <= 0) || (curDiff < 0 && prevDiff >= 0) {
			res++
			prevDiff = curDiff
		}
	}

	return res
}

func wiggleMaxLength2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	up := 1
	down := 1
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 {
			up = down + 1
		} else if diff < 0 {
			down = up + 1
		}
	}

	if up > down {
		return up
	}
	return down
}
