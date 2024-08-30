package largestrectangleinhistogram

func largestRectangleArea(heights []int) int {
	res := 0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack := []int{}

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				w := i - left - 1
				res = max(heights[mid]*w, res)
			}
		}
		stack = append(stack, i)
	}

	return res
}
