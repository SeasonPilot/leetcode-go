package main

//链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters/solution/jian-dan-yi-dong-javac-pythonjshua-dong-bff20/

// 3. 优化后的滑动窗口
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func lengthOfLongestSubstring3(s string) int {
	var n = len(s)
	if n <= 1 {
		return n
	}

	var maxLen = 1
	var left, right, window = 0, 0, make(map[byte]int) // map 存储字符和出现的index

	for right < n {
		var rightChar = s[right]
		var rightCharIndex = 0

		// 有重复字符
		if _, ok := window[rightChar]; ok { // rightChar 是否出现过
			rightCharIndex = window[rightChar] // 拿到重复字符之前出现的 index
		}
		if rightCharIndex > left {
			left = rightCharIndex // 缩小左边界
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1 // 更新 result 值
		}

		// 无重复字符
		window[rightChar] = right + 1 // rightChar 没有出现过，记录 rightChar 和 index
		right++                       // 移动 右边界
	}

	return maxLen
}
