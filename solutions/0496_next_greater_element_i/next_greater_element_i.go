package nextgreaterelementi

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = -1
	}
	m := make(map[int]int)
	for i, v := range nums1 {
		m[v] = i
	}
	stack := []int{}

	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			if idx, ok := m[stack[len(stack)-1]]; ok {
				res[idx] = nums2[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	return res
}
