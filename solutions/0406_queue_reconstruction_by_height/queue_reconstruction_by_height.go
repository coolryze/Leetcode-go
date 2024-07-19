package queuereconstructionbyheight

import (
	"slices"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	res := [][]int{}
	for _, p := range people {
		res = slices.Insert(res, p[1], p)
	}
	return res
}
