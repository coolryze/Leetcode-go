package subsetsii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))

	var backtracking func(nums []int, start int)
	backtracking = func(nums []int, start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtracking(nums, i+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	sort.Ints(nums)
	backtracking(nums, 0)

	return res
}
