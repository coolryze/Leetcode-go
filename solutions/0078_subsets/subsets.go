package subsets

func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtracking func(nums []int, start int)
	backtracking = func(nums []int, start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(nums, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(nums, 0)

	return res
}
