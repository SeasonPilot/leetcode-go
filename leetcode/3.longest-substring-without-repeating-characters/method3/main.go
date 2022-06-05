package main

// 最终使用此方法
// leetcode/0003.Longest-Substring-Without-Repeating-Characters/3. Longest Substring Without Repeating Characters.go
// 解法三 滑动窗口-哈希桶
func lengthOfLongestSubstring2(s string) int {
	left, right, res := 0, 0, 0           // 左右边界
	indexes := make(map[byte]int, len(s)) // 哈希集合，记录每个字符是否出现过  // map 存储 字符 和 index

	for right < len(s) {
		// 有重复字符，先缩小左边界
		if idx, ok := indexes[s[right]]; ok && idx >= left { // 右边界字符 rightChar 之前出现过并且 位置>=左边界。  因为是要缩小左边界，所以 idx 要>= left
			left = idx + 1 // 缩小左边界。 移动左边界到 rightChar 之前出现过的 index + 1 位置
		}
		// 无重复字符
		indexes[s[right]] = right // 更新 rightChar 的 index 或 存储当前字符的 index 信息。
		right++                   // 移动右边界

		res = max(res, right-left) // 更新最大长度。 因为先移动了右边界，所以这里不用 +1
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。一旦出现了重复字符，就需要缩小左边界，直到重复的字符移出了左边界，
// 然后继续移动滑动窗口的右边界。以此类推，每次移动需要计算当前长度，并判断是否需要更新最大长度，最终最大的值就是题目中的所求。
