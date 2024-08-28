package nextgreaterelementii

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
	stack := []int{}

	for i := 0; i < len(nums)*2; i++ {
		for len(stack) > 0 && nums[i%len(nums)] > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = nums[i%len(nums)]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%len(nums))
	}

	return res
}
