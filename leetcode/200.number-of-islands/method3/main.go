package main

// flood fill   DFS
func numIslands(grid [][]byte) int {
	var islandsNums int
	for i := 0; i < len(grid); i++ { // i 是第一个 index
		for j := 0; j < len(grid[i]); j++ { // j 是第二个 index  fixme: len(grid[i]) 方括号内为 i,不是 0
			if grid[i][j] == '1' {
				islandsNums += dfs(grid, i, j)
			}
		}
	}

	return islandsNums
}

func dfs(grid [][]byte, i, j int) int { // i,j 分别表示 第一个 index 和 第二个 index
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	// terminator
	if grid[i][j] == '0' {
		return 0
	}

	// process logic in current level
	grid[i][j] = '0'
	for k := 0; k < len(dx); k++ {
		x, y := i+dx[k], j+dy[k] // x 是第一个 index , y 是第二个 index

		if x < len(grid) && x >= 0 && y < len(grid[i]) && y >= 0 { // 判断是否超出边界   // fixme: grid 行和列长度不同
			if grid[x][y] == '1' {
				// drill down
				dfs(grid, x, y)
			}
		}
	}

	return 1
}
