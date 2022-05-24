package main

// flood fill
func numIslands(grid [][]byte) int {
	var islandsNums int
	//for i := 0; i < len(grid[0]); i++ {
	//	for j := 0; j < len(grid); j++ {
	//
	//	}
	//}
	for i, row := range grid { // fixme: j,i 搞反了，列对应x，行对应y
		for j, b := range row {
			if b == '1' {
				islandsNums += dfs(grid, i, j)
			}
		}
	}

	return islandsNums
}

func dfs(grid [][]byte, i, j int) int {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	// terminator
	if grid[i][j] == '0' { // fixme: grid[][]  第一个 index 对应行,行对应 y ; 列对应 x
		return 0
	}

	// process logic in current level
	grid[i][j] = '0'
	for k := 0; k < len(dx); k++ {
		x, y := j+dx[k], i+dy[k]

		if x < len(grid[0]) && x >= 0 && y < len(grid) && y >= 0 { // 判断是否超出边界   // fixme: grid 行和列长度不同
			if grid[y][x] == '1' { // fixme: 这里是 x
				// drill down
				dfs(grid, y, x)
			}
		}
	}

	return 1
}
