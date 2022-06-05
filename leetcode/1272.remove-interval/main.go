package main

// https://www.codeleading.com/article/32953238819/
func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	res := make([][]int, 0)
	for _, inter := range intervals {
		if inter[0] >= toBeRemoved[1] || inter[1] <= toBeRemoved[0] { // 不相交
			res = append(res, inter)
		} else { // 相交
			// 先判断头
			if inter[0] < toBeRemoved[0] {
				// 在相交的条件下，toBeRemoved[0] 一定是 < inter[1]的。
				// 如果 inter[0] > toBeRemoved[0] 则 相交部分都被 remove 了，不需要添加到结果中了。
				res = append(res, []int{inter[0], toBeRemoved[0]})
			}
			// 再判断尾
			if inter[1] > toBeRemoved[1] {
				res = append(res, []int{toBeRemoved[1], inter[1]})
			}
		}
	}
	return res
}

// 错误
func practice1(intervals [][]int, toBeRemoved []int) [][]int {
	var res [][]int
	for _, interval := range intervals {
		if interval[0] > toBeRemoved[1] || interval[1] < toBeRemoved[0] { // 不相交
			res = append(res, interval)
		} else { // 相交
			if interval[0] < toBeRemoved[1] {
				res = append(res, []int{toBeRemoved[1], interval[1]})
			}
			if interval[1] > toBeRemoved[0] {
				res = append(res, []int{interval[0], toBeRemoved[0]})
			}
		}
	}
	return res
}

// 忘记 == 情况了
func practice2(intervals [][]int, toBeRemoved []int) [][]int {
	var res [][]int
	for _, inter := range intervals {
		if inter[0] >= toBeRemoved[1] || inter[1] <= toBeRemoved[0] { // 忘记 == 情况了
			res = append(res, inter)
		} else {
			if inter[0] < toBeRemoved[0] {
				res = append(res, []int{inter[0], toBeRemoved[0]})
			}
			if inter[1] > toBeRemoved[1] {
				res = append(res, []int{toBeRemoved[1], inter[1]})
			}
		}
	}
	return res
}
