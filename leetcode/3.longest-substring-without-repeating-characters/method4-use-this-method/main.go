package main

// 最终使用此方法
// 滑动窗口 labuladong。这个方法更容易记，且使用 滑动窗口算法框架,https://labuladong.github.io/article/fname.html?fname=滑动窗口技巧进阶
func lengthOfLongestSubstring(s string) int {
	windows := make(map[byte]int)

	left, right, ret := 0, 0, 0

	for right < len(s) {
		c := s[right]
		right++
		windows[c]++ // 进行窗口内数据的一系列更新

		for windows[c] > 1 { // 判断左侧窗口是否要收缩
			d := s[left]
			left++
			windows[d]-- // 进行窗口内数据的一系列更新
		}
		// 在这里更新答案
		ret = max(right-left, ret)
	}
	return ret
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
