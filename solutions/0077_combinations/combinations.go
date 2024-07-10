package combinations

func combine(n int, k int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtracking func(int, int, int)
	backtracking = func(n, k, start int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}

		for i := start; i <= n; i++ {
			if n-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			backtracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(n, k, 1)

	return res
}
