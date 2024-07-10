package combinationsumiii

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtracking func(int, int, int, int)
	backtracking = func(k, n, start, sum int) {
		if len(path) == k {
			if sum == n {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}

		for i := start; i <= 9; i++ {
			if sum+i > n || 9-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			backtracking(k, n, i+1, sum+i)
			path = path[:len(path)-1]
		}
	}

	backtracking(k, n, 1, 0)

	return res
}
