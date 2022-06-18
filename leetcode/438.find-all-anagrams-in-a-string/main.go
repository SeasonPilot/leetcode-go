package main

// 滑动窗口
func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for _, c := range p {
		need[byte(c)]++
	}

	left, right := 0, 0
	valid := 0
	var ret []int

	for right < len(s) {
		c := s[right]
		right++
		if v, ok := need[c]; ok {
			window[c]++         // 先计数
			if window[c] == v { // 是否满足条件
				valid++
			}
		}

		for right-left >= len(p) { // 23-26行，需要改变的地方
			if valid == len(need) {
				ret = append(ret, left)
			}

			d := s[left]
			left++
			if v, ok := need[d]; ok {
				if window[d] == v {
					valid--
				}
				window[d]--
			}

		}
	}
	return ret
}
