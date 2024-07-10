package lettercombinationsofaphonenumber

func letterCombinations(digits string) []string {
	res := []string{}
	if digits == "" {
		return res
	}

	m := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path := []byte{}

	var backtracking func(string, int)
	backtracking = func(digits string, start int) {
		if len(path) == len(digits) {
			res = append(res, string(path))
			return
		}

		digit := int(digits[start] - '0')
		str := m[digit]
		for i := 0; i < len(str); i++ {
			path = append(path, str[i])
			backtracking(digits, start+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(digits, 0)

	return res
}
