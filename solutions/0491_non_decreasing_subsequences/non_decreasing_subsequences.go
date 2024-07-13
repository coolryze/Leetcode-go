package nondecreasingsubsequences

func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtracking func(nums []int, start int)
	backtracking = func(nums []int, start int) {
		if len(path) > 1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		used := make(map[int]bool, len(nums))
		for i := start; i < len(nums); i++ {
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			if used[nums[i]] {
				continue
			}

			path = append(path, nums[i])
			used[nums[i]] = true
			backtracking(nums, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(nums, 0)

	return res
}
