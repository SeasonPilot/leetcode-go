package main

import "math"

func minWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for _, c := range t {
		need[byte(c)]++
	}

	left, right := 0, 0
	valid := 0
	length := math.MaxInt32
	start := 0

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) { // fixme: 嵌套 for 循环
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
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
