package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n) // 记录每根棒子的左右边界
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{}    // 存储的是 heights 的 index
	for i := 0; i < n; i++ { // 遍历所有棒子
		// 当前棒子高度 <= 栈顶元素值; 证明栈顶棒子的右边界找到了。  for 循环，一直弹栈到 当前棒子高度 > 栈顶元素值 为止。
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			// 当前栈顶元素(棒子)  的右边界就是 i(当前循环的棒子)；记录的是 栈顶棒子 的右边界
			right[mono_stack[len(mono_stack)-1]] = i
			// 弹出栈顶 后的新栈
			mono_stack = mono_stack[:len(mono_stack)-1]
		}

		// 当前棒子高度 > 栈顶元素值
		if len(mono_stack) == 0 { // 初始的栈为空，设置第一根棒子的左边界为-1，哨兵
			left[i] = -1
		} else {
			// 当前棒子的左边界 就是 栈顶的元素(或新栈的栈顶元素)；即 记录第 i 个棒子的左边界
			left[i] = mono_stack[len(mono_stack)-1]
		}
		// 压栈；当前棒子高度 > 栈顶元素值
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ { //  遍历所有棒子，算出面积
		ans = max(
			ans,
			(right[i]-left[i]-1)*heights[i],
		)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
