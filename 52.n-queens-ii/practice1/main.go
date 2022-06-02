package main

// 基于位运算的回溯
func totalNQueens(n int) (ans int) {
	var solve = func(rows, columns, diagonals1, diagonals2 int) {}
	solve = func(rows, columns, diagonals1, diagonals2 int) {
		// terminator
		if rows == n {
			ans++
			return
		}

		// process logic in current level
		availablePosition := ^(columns | diagonals1 | diagonals2) & (1<<n - 1) // fixme: (1<<n - 1) 要加括号

		for availablePosition > 0 { // fixme: 应该是 availablePosition > 0
			p := availablePosition & -availablePosition
			availablePosition &= availablePosition - 1
			// drill down
			solve(rows+1, columns|p, (diagonals1|p)<<1, (diagonals2|p)>>1) // fixme: << 优先级大于 |
		}
		// reverse
	}

	solve(0, 0, 0, 0)
	return
}
