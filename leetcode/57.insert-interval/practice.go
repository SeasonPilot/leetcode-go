package main

func insert1(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	var ans [][]int
	var merge bool

	for _, interval := range intervals {
		if interval[0] > right {
			if !merge {
				ans = append(ans, []int{left, right})
				merge = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			ans = append(ans, interval)
		} else {
			left = min(interval[0], left)
			right = max(interval[1], right)
		}
	}
	if !merge {
		ans = append(ans, []int{left, right})
	}

	return ans
}
