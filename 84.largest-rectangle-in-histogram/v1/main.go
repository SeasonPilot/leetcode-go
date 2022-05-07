package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
}

func largestRectangleArea(heights []int) int {
	//单调栈（单调递增）
	stack := make([]int, 0)
	stack = append(stack, -1)    //stack的哨兵，方便确定左边界
	heights = append(heights, 0) //添加一个哨兵，减少代码量
	ln := len(heights)
	res := 0 //结果

	for i := 0; i < ln; i++ {
		//因为我们无法访问heights[-1]，所以限制len(stack) > 1
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			//栈顶元素，也就是当前要求的矩形柱子的下标
			top := stack[len(stack)-1]
			//出栈
			stack = stack[:len(stack)-1]
			//左边界（栈顶元素的后一个元素）
			l := stack[len(stack)-1]
			//矩形面积：(右边界-左边界-1) * 高度
			//右边界就是i
			//高度就是以栈顶元素为下标的柱子的高度
			//左边界就是栈顶元素的下一位元素（因为我们添加了哨兵-1，所以这公式依旧成立）
			res = max(res, (i-l-1)*heights[top])
		}

		stack = append(stack, i)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
