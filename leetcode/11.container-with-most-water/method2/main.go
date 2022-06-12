package main

// https://labuladong.github.io/algo/4/31/128/#扩展延伸
// 双指针
func maxArea(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	left, right := 0, n-1 // 左右指针
	ret := 0
	for left < right {
		if height[left] < height[right] { // 双指针技巧，移动较低的一边
			ret = max(ret, (right-left)*height[left])
			left++
		} else {
			ret = max(ret, (right-left)*height[right])
			right--
		}
	}
	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

/*
	其实也好理解，因为矩形的高度是由 min(height[left], height[right]) 即较低的一边决定的：
	你如果移动较低的那一边，那条边可能会变高，使得矩形的高度变大，进而就「有可能」使得矩形的面积变大；
	相反，如果你去移动较高的那一边，矩形的高度是无论如何都不会变大的，所以不可能使矩形的面积变得更大。
*/
