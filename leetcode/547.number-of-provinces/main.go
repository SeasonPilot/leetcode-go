package main

// leetCode 实现方式
func findCircleNum(isConnected [][]int) int {
	// init define
	n := len(isConnected)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	// [0,1,2]

	// find define
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// union define
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	// 按行遍历二维数组
	for i, row := range isConnected {
		for j := i + 1; j < n; j++ {
			if row[j] == 1 { // 如果值为 1，则合并
				union(i, j)
			}
		}
	}

	var ans int
	for i, p := range parent {
		if i == p {
			ans++
		}
	}
	return ans
}
