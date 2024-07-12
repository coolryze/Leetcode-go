package combinationsumii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	used := make([]bool, len(candidates))

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
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
				continue
			}
			path = append(path, candidates[i])
			used[i] = true
			backtracking(candidates, target, i+1, sum+candidates[i])
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	sort.Ints(candidates)
	backtracking(candidates, target, 0, 0)

	return res
}
