package permutationsii

import "sort"

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))

	var backtracking func(nums []int)
	backtracking = func(nums []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		usedMap := make(map[int]bool, len(nums))
		for i := 0; i < len(nums); i++ {
			if used[i] || usedMap[nums[i]] {
				continue
			}

			usedMap[nums[i]] = true
			used[i] = true
			path = append(path, nums[i])
			backtracking(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtracking(nums)

	return res
}

func permuteUnique2(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))

	var backtracking func(nums []int)
	backtracking = func(nums []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backtracking(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	sort.Ints(nums)
	backtracking(nums)

	return res
}
