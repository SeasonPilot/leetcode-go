package main

// 最终使用此方法
// flood fill   DFS
func numIslands(grid [][]byte) int {
	var islandsNums int
	for i := 0; i < len(grid); i++ { // 第一个 index
		for j := 0; j < len(grid[0]); j++ { // 第二个 index
			if grid[i][j] == '1' {
				islandsNums += dfs(grid, i, j)
			}
		}
	}

	return islandsNums
}

func dfs(grid [][]byte, i, j int) int { // i,j 分别表示 第一个 index 和 第二个 index
	di := []int{-1, 1, 0, 0}
	dj := []int{0, 0, -1, 1}

	// terminator
	if grid[i][j] == '0' { // fixme: grid[][]  第一个 index 对应行 y ; 第二个 index 对应列 x
		return 0
	}

	// process logic in current level
	grid[i][j] = '0'
	for k := 0; k < len(di); k++ {
		ni, nj := i+di[k], j+dj[k]

		if ni < len(grid) && ni >= 0 && nj < len(grid[0]) && nj >= 0 { // 判断是否超出边界   // fixme: grid 行和列长度不同
			if grid[ni][nj] == '1' { // fixme: 这里是 ni
				// drill down
				dfs(grid, ni, nj)
			}
		}
	}

	return 1
}
