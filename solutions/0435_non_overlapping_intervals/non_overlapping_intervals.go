package nonoverlappingintervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	res := 0

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
			res++
		}
	}

	return res
}

func eraseOverlapIntervals2(intervals [][]int) int {
	res := 1

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i-1][1] {
			res++
		} else {
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
		}
	}

	return len(intervals) - res
}
