package main

// 滑动窗口
func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	for _, c := range s1 {
		need[byte(c)]++
	}

	left, right := 0, 0
	valid := 0

	for right < len(s2) {
		c := s2[right]
		right++
		if v, ok := need[c]; ok {
			window[c]++
			if window[c] == v {
				valid++
			}
		}

		for right-left >= len(s1) { // 2、什么条件下，窗口应该暂停扩大，开始移动 left 缩小窗口？  22-25行，需要改变的地方
			if valid == len(need) { // 4、我们要的结果应该在扩大窗口时还是缩小窗口时进行更新？
				return true
			}

			d := s2[left]
			left++
			if v, ok := need[d]; ok {
				if window[d] == v {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
