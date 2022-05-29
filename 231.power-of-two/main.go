package main

// 一个数 n 是 2 的幂，当且仅当 n 是正整数，并且 n 的二进制表示中仅包含 1 个 1。
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// n&(n-1) == 0 打掉最低位的 1 后应该为 0
