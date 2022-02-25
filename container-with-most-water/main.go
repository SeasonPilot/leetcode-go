package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	var area int
	// i j 是 index, 不是数组中的值
	for i := 0; i < len(height); i++ {
		// j 初始值为 1, j 最终是移动到最后一个索引的位置,j 为索引下标
		for j := 1; j < len(height); j++ {
			// 应该是 宽度(j - i),不是项数 (j - i + 1)
			tmp := (j - i) * Min(height[i], height[j])
			if tmp > area {
				area = tmp
			}
		}
	}
	return area
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
