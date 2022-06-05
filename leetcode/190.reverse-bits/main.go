package main

// 位运算
func reverseBits(n uint32) (rev uint32) {
	for i := 0; i < 32 && n > 0; i++ { // 从低位往高位枚举 n 的每一位
		rev |= (n & 1) << (31 - i) // 这里 (n & 1) 表示 最后一位是 1 还是 0。 把最后一位左移到对应的位置上，并和 rev 进行按位或,相当于和对应位置的 0 按位或。  // 合并结果
		n >>= 1                    // 每枚举一位就将 n 右移一位，这样当前 n 的最低位就是我们要枚举的比特位。
	}
	return
}

//链接：https://leetcode.cn/problems/reverse-bits/solution/dian-dao-er-jin-zhi-wei-by-leetcode-solu-yhxz/
/*
	判断奇偶：
	x % 2 == 1 —> (x & 1) == 1
	x % 2 == 0 —> (x & 1) == 0
*/
