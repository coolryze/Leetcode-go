package repeatedsubstringpattern

func repeatedSubstringPattern(s string) bool {
	l := len(s)
	next := make([]int, l)
	j := -1
	next[0] = j
	for i := 1; i < l; i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}

	if next[l-1] != -1 && l%(l-(next[l-1]+1)) == 0 {
		return true
	}
	return false
}
