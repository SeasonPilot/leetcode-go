package main

func findCircleNum1(isConnected [][]int) (ans int) {
	// init define
	n := len(isConnected)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	// find define
	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			x = parent[x]
		}
		return x
	}

	// union define
	// 0,1
	union := func(from, to int) {
		parent[find(from)] = find(to) // parent[0] = 1
	}

	// 按行遍历二维数组
	for i, row := range isConnected {
		for j := i + 1; j < n; j++ {
			if row[j] == 1 {
				union(i, j)
			}
		}
	}

	for i, p := range parent {
		if i == p {
			ans++
		}
	}
	return
}
