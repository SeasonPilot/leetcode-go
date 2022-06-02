package main

// 基于位运算的回溯  最终使用此方法
func totalNQueens(n int) (ans int) {
	var solve func(row, columns, diagonals1, diagonals2 int)

	solve = func(row, columns, diagonals1, diagonals2 int) { // columns, diagonals1, diagonals2 表示不能放的位置
		// terminator
		if row == n {
			ans++
			return
		}

		// process logic in current level
		availablePositions := ^(columns | diagonals1 | diagonals2) & (1<<n - 1) // 得到当前所有的空位   1 表示可以放置皇后。  1<<n - 1 表示 n 个 1 , 是用来限制二进制位数的？即每行有多少格子？

		for availablePositions > 0 {
			p := availablePositions & -availablePositions // 得到最低位的 1  放入皇后的位置
			availablePositions &= availablePositions - 1  // 清零最低位的 1  表示在 p 位置上放⼊皇后。  放置皇后后相应位置变为 0,即表示该位置不能再放置皇后，
			// drill down
			solve(row+1, columns|p, (diagonals1|p)<<1, (diagonals2|p)>>1)
		}
	}

	solve(0, 0, 0, 0)
	return
}
