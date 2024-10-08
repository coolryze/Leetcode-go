package findtheindexofthefirstoccurrenceinastring

// kmp: next
func strStr(haystack string, needle string) int {
	next := make([]int, len(needle))
	j := 0
	next[0] = j
	for i := 1; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}

	j = 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}

	return -1
}

// kmp: next-1
func strStr2(haystack string, needle string) int {
	next := make([]int, len(needle))
	j := -1
	next[0] = j
	for i := 1; i < len(needle); i++ {
		for j >= 0 && needle[i] != needle[j+1] {
			j = next[j]
		}
		if needle[i] == needle[j+1] {
			j++
		}
		next[i] = j
	}

	j = -1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}

	return -1
}

// 暴力
func strStr3(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		if i > len(haystack)-len(needle) {
			break
		}
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1
}
