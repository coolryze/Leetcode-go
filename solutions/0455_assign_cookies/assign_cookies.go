package assigncookies

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	gi := 0
	for i := 0; i < len(s); i++ {
		if gi < len(g) && s[i] >= g[gi] {
			gi++
		}
	}

	return gi
}

func findContentChildren2(g []int, s []int) int {
	res := 0
	sort.Ints(g)
	sort.Ints(s)

	si := len(s) - 1
	for i := len(g) - 1; i >= 0; i-- {
		if si >= 0 && s[si] >= g[i] {
			res++
			si--
		}
	}

	return res
}
