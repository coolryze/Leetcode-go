package partitionlabels

func partitionLabels(s string) []int {
	res := []int{}

	m := [26]int{}
	for i, c := range s {
		m[c-'a'] = i
	}

	left := 0
	right := 0
	for i := 0; i < len(s); i++ {
		right = max(m[s[i]-'a'], right)
		if right == i {
			res = append(res, right-left+1)
			left = right + 1
		}
	}

	return res
}
