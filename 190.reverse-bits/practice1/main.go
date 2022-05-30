package main

func reverseBits(num uint32) (rev uint32) {
	for i := 0; i < 32 && num > 0; i++ {
		r := (num & 1) << (31 - i) // fixme: 这里不是 32
		rev |= r                   // 合并结果
		num = num >> 1
	}
	return
}

func reverseBits1(num uint32) (rev uint32) {
	for i := 0; i < 32 && num > 0; i++ {
		rev |= (num & 1) << (31 - i) // fixme: 这里不是 32
		num >>= 1
	}
	return
}
