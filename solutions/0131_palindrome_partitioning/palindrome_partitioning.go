package palindromepartitioning

func partition(s string) [][]string {
	res := [][]string{}
	path := []string{}

	isPalindrome := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}

		return true
	}

	var backtracking func(s string, start int)
	backtracking = func(s string, start int) {
		if start == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(s); i++ {
			str := s[start : i+1]
			if !isPalindrome(str) {
				continue
			}

			path = append(path, str)
			backtracking(s, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(s, 0)

	return res
}
