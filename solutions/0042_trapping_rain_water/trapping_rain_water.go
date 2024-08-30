package trappingrainwater

func trap(height []int) int {
	res := 0

	stack := []int{}

	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				h := min(height[i], height[left]) - height[mid]
				w := i - left - 1
				res += h * w
			}
		}
		stack = append(stack, i)
	}

	return res
}

func trap2(height []int) int {
	res := 0

	leftMax := make([]int, len(height))
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}

	rightMax := make([]int, len(height))
	rightMax[len(rightMax)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}

	for i := 0; i < len(height)-1; i++ {
		count := min(leftMax[i], rightMax[i]) - height[i]
		if count > 0 {
			res += count
		}
	}

	return res
}
