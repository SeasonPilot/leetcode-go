package main

func hammingWeight(num uint32) int {
	var sum int
	for num > 0 {
		num &= num - 1
		sum++
	}
	return sum
}

// 把所有的 1 都清掉了，最后就是 0，所以退出条件就是等于 0

func hammingWeight2(num uint32) int {
	var sum int
	for ; num > 0; num &= num - 1 {
		sum++
	}
	return sum
}

func hammingWeight3(num uint32) (ones int) {
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return
}
