package practice2

func lengthOfLongestSubstring(s string) int {
	windows := make(map[byte]int)

	left, right := 0, 0
	ret := 0
	for right < len(s) {
		c := s[right]
		windows[c]++
		right++

		for windows[c] > 1 { // 判断左侧窗口是否要收缩
			d := s[left]
			windows[d]--
			left++
		}

		ret = max(right-left, ret)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
