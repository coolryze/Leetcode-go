package palindromicsubstrings

func countSubstrings(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	n := len(s)

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] {
					res++
				}
			}
		}
	}

	return res
}

func countSubstrings2(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	n := len(s)

	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				if l == 1 || l == 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] {
					res++
				}
			}
		}
	}

	return res
}

func countSubstrings3(s string) int {
	res := 0
	bytes := []byte(s)

	for i := 0; i < len(bytes); i++ {
		for j := i + 1; j <= len(bytes); j++ {
			if isPalindrome(bytes[i:j]) {
				res += 1
			}
		}
	}

	return res
}

func isPalindrome(bytes []byte) bool {
	for i := 0; i <= len(bytes)/2; i++ {
		if bytes[i] != bytes[len(bytes)-i-1] {
			return false
		}
	}
	return true
}
