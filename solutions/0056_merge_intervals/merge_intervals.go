package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	res := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > right {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = max(intervals[i][1], right)
		}
	}
	res = append(res, []int{left, right})

	return res
}

func merge2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	removed := make([]bool, len(intervals))
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			intervals[i][0] = intervals[i-1][0]
			intervals[i][1] = max(intervals[i][1], intervals[i-1][1])
			removed[i-1] = true
		}
	}

	res := [][]int{}
	for i, val := range removed {
		if !val {
			res = append(res, intervals[i])
		}
	}

	return res
}
