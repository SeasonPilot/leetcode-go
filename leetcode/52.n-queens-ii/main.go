package main

//链接：https://leetcode.cn/problems/n-queens-ii/solution/nhuang-hou-ii-by-leetcode-solution/
func totalNQueens(n int) (ans int) {
	var solve func(row, columns, diagonals1, diagonals2 int)

	solve = func(row, columns, diagonals1, diagonals2 int) {
		if row == n {
			ans++
			return
		}

		availablePositions := (1<<n - 1) &^ (columns | diagonals1 | diagonals2) // 如果运算符右侧数值的第 i 位为 1，那么计算结果中的第 i 位为 0；如果运算符右侧数值的第 i 位为 0，那么计算结果中的第 i 位为运算符左侧数值的第 i 位的值。

		for availablePositions > 0 {
			position := availablePositions & -availablePositions //  得到最低位的 1
			solve(row+1, columns|position, (diagonals1|position)<<1, (diagonals2|position)>>1)
			availablePositions &^= position // 移除该比特位
		}
	}

	solve(0, 0, 0, 0)
	return
}
