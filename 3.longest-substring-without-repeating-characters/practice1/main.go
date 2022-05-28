package main

func lengthOfLongestSubstring(s string) int {
	left, right, ret := 0, 0, 0
	indexes := make(map[byte]int)

	for right < len(s) {
		if idx, ok := indexes[s[right]]; ok && idx >= left {
			left = idx + 1
		}

		indexes[s[right]] = right
		right++

		ret = max(ret, right-left)
	}

	return ret
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
