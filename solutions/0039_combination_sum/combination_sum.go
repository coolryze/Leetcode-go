package combinationsum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtracking func(candidates []int, target, start, sum int)
	backtracking = func(candidates []int, target, start, sum int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			backtracking(candidates, target, i, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}

	sort.Ints(candidates)
	backtracking(candidates, target, 0, 0)

	return res
}
