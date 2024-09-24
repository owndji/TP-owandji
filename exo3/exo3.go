package exo3

import (
	"sort"
)

func Ft_non_overlap(intervals [][]int) int {

	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	remove_count := 0

	last_end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < last_end {
			remove_count++
		} else {
			last_end = intervals[i][1]
		}
	}

	return remove_count
}
