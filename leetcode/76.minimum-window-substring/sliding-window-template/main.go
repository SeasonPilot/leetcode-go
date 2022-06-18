package main

import "fmt"

// 滑动窗口模板
func minWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for _, c := range t {
		need[byte(c)]++
	}

	left, right := 0, 0
	valid := 0

	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 增大窗口
		right++
		// 进行窗口内数据的一系列更新
		...

		/*** debug 输出的位置 ***/
		fmt.Printf("window: [%d, %d)\n", left, right)
		/********************/

		// 判断左侧窗口是否要收缩
		for (window needs shrink) {
			// 在这里更新最小覆盖子串
			...
			// d 是将移出窗口的字符
			d := s[left]
			// 缩小窗口
			left++
			// 进行窗口内数据的一系列更新
			...
		}
	}

}

/*
	现在开始套模板，只需要思考以下四个问题：

	1、当移动 right 扩大窗口，即加入字符时，应该更新哪些数据？

	2、什么条件下，窗口应该暂停扩大，开始移动 left 缩小窗口？

	3、当移动 left 缩小窗口，即移出字符时，应该更新哪些数据？

	4、我们要的结果应该在扩大窗口时还是缩小窗口时进行更新？
*/
