package maximumsubarray

import "math"

func maxSubArray(nums []int) int {
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
