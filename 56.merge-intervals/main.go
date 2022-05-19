package main

import "sort"

// https://leetcode.cn/problems/merge-intervals/solution/shou-hua-tu-jie-56he-bing-qu-jian-by-xiao_ben_zhu/
func merge(intervals [][]int) [][]int {
	// 1. 按照区间左边的值进行升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 初始化 res, 用于存储合并之后区间结果的动态数组
	res := [][]int{}
	prev := intervals[0]
	// 3. 遍历处理每一个区间
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 没有一点重合
			res = append(res, prev) // 直接放入结果集
			prev = cur
		} else { // 有重合
			prev[1] = max(prev[1], cur[1]) // 将 prev 扩大，更新右边界的值
		}
	}
	res = append(res, prev)
	return res
}

// https://leetcode.cn/problems/merge-intervals/solution/shi-pin-wen-zi-xiang-xi-jiang-jie-he-bing-qu-jian-/
func merge1(intervals [][]int) [][]int {
	// 按照 end 升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var outputs = [][]int{}
	outputs = append(outputs, intervals[0])
	for i := 1; i < len(intervals); i++ {
		var lastInterval = outputs[len(outputs)-1]
		var currLeft = intervals[i][0]
		var currRight = intervals[i][1]
		if lastInterval[1] < currLeft {
			outputs = append(outputs, intervals[i])
		} else {
			lastInterval[1] = max(lastInterval[1], currRight)
		}
	}
	return outputs
}

// 使用老汤的思路，在笨猪爆破组代码基础上进行修改的版本。  最优！！！
func merge2(intervals [][]int) [][]int {
	// 1. 按照区间左边的值进行升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 初始化 res, 用于存储合并之后区间结果的动态数组
	res := [][]int{}
	res = append(res, intervals[0])
	// 3. 遍历处理每一个区间
	for i := 1; i < len(intervals); i++ { // for range 不能指定开始循环的位置
		last := res[len(res)-1] // 结果中的最后一个元素
		cur := intervals[i]
		if last[1] < cur[0] { // 没有一点重合
			res = append(res, cur) // 直接放入结果集
		} else { // 有重合
			last[1] = max(last[1], cur[1]) // 将 last 扩大，改变右边界的值
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func practice1(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ { // for range 不能指定开始循环的位置
		last := res[len(res)-1]
		cur := intervals[i]
		if last[1] < cur[0] {
			res = append(res, cur)
		} else {
			last[1] = max(last[1], cur[1])
		}
	}
	return res
}
