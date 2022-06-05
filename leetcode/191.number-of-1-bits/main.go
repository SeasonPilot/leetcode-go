package main

import "math/bits"

// 最终使用此方法
// 位运算优化
// 链接：https: //leetcode.cn/problems/number-of-1-bits/solution/wei-1de-ge-shu-by-leetcode-solution-jnwf/
func hammingWeight(num uint32) (ones int) {
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return
}

// 解法二
func hammingWeight2(num uint32) int {
	return bits.OnesCount(uint(num))
}
