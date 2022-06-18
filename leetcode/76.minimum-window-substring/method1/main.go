package main

import (
	"fmt"
	"math"
)

// 最终使用此方法
// https://labuladong.github.io/algo/2/18/25/
func minWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int) // 计数器
	for _, c := range t {
		need[byte(c)]++
	}

	left, right := 0, 0
	valid := 0 // valid 变量表示窗口中满足 need 条件的字符个数
	// 记录最小覆盖子串的起始索引及长度
	start, length := 0, math.MaxInt32

	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 增大窗口
		right++
		// 进行窗口内数据的一系列更新
		if v, ok := need[c]; ok {
			window[c]++
			if window[c] == v {
				valid++
			}
		}

		/*** debug 输出的位置 ***/
		fmt.Printf("window: [%d, %d)\n", left, right)
		/********************/

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}

			// d 是将移出窗口的字符
			d := s[left]
			// 缩小窗口
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] { //
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
