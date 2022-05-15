package main

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false // newInterval 是否有合入

	for _, interval := range intervals {
		if interval[0] > right {
			// interval 在插入区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			// interval 在插入区间的左侧且无交集
			ans = append(ans, interval)
		} else {
			// interval 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return
}

//那么我们应当在什么时候将区间 S 加入答案呢？由于我们需要保证答案也是按照左端点排序的，因此当我们遇到第一个 满足 li > right
//的区间时，说明以后遍历到的区间不会与 S 重叠，并且它们左端点一定会大于 S 的左端点。此时我们就可以将 S 加入答案。
//特别地，如果不存在这样的区间，我们需要在遍历结束后，将 S 加入答案。

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
