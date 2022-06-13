package main

// 双指针
func trap(height []int) int {
	n := len(height)
	//left, right := 1, n-2
	//lMax, rMax := height[0], height[n-1]

	left, right := 0, n-1 // fixme: 从左右端点开始移动不会漏掉
	lMax, rMax := 0, 0

	ret := 0

	for left < right {
		lMax = max(height[left], lMax)
		rMax = max(height[right], rMax)
		if lMax < rMax { // 只需要比较 lMax 和 rMax 即可。 其实这个问题要这么思考，我们只在乎 min(l_max, r_max)。对于上图的情况，我们已经知道 l_max < r_max 了，至于这个 r_max 是不是右边最大的，不重要。重要的是 height[i] 能够装的水只和较低的 l_max 之差有关：
			ret += lMax - height[left]
			left++
		} else {
			ret += rMax - height[right]
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
https://labuladong.github.io/algo/4/31/128/#:~:text=%E5%85%B6%E5%AE%9E%E8%BF%99%E4%B8%AA%E9%97%AE%E9%A2%98%E8%A6%81%E8%BF%99%E4%B9%88%E6%80%9D%E8%80%83%EF%BC%8C%E6%88%91%E4%BB%AC%E5%8F%AA%E5%9C%A8%E4%B9%8E%20min(l_max%2C%20r_max)%E3%80%82%E5%AF%B9%E4%BA%8E%E4%B8%8A%E5%9B%BE%E7%9A%84%E6%83%85%E5%86%B5%EF%BC%8C%E6%88%91%E4%BB%AC%E5%B7%B2%E7%BB%8F%E7%9F%A5%E9%81%93%20l_max%20%3C%20r_max%20%E4%BA%86%EF%BC%8C%E8%87%B3%E4%BA%8E%E8%BF%99%E4%B8%AA%20r_max%20%E6%98%AF%E4%B8%8D%E6%98%AF%E5%8F%B3%E8%BE%B9%E6%9C%80%E5%A4%A7%E7%9A%84%EF%BC%8C%E4%B8%8D%E9%87%8D%E8%A6%81%E3%80%82%E9%87%8D%E8%A6%81%E7%9A%84%E6%98%AF%20height%5Bi%5D%20%E8%83%BD%E5%A4%9F%E8%A3%85%E7%9A%84%E6%B0%B4%E5%8F%AA%E5%92%8C%E8%BE%83%E4%BD%8E%E7%9A%84%20l_max%20%E4%B9%8B%E5%B7%AE%E6%9C%89%E5%85%B3%EF%BC%9A
*/
