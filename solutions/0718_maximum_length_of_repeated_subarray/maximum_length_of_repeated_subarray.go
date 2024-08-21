package maximumlengthofrepeatedsubarray

func findLength(nums1 []int, nums2 []int) int {
	res := 0
	dp := make([]int, len(nums1)+1)

	for i := 1; i <= len(nums1); i++ {
		for j := len(nums2); j > 0; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0
			}
			if dp[j] > res {
				res = dp[j]
			}
		}
	}

	return res
}

func findLength2(nums1 []int, nums2 []int) int {
	res := 0
	dp := make([][]int, len(nums1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}
