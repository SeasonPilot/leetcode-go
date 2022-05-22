package main

// 双向 BFS
// 练习 127.word-ladder/method3java.go
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 1. 将 wordList 转换成 hash
	wordSet := make(map[string]struct{})
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}
	// 边界判断
	if _, has := wordSet[endWord]; !has {
		return 0
	}

	// 2. 将 beginWord endWord 添加到 beginSet, endSet hash
	beginSet, endSet := make(map[string]struct{}), make(map[string]struct{})
	beginSet[beginWord] = struct{}{}
	endSet[endWord] = struct{}{}

	// 3. 记录访问过的 BFS 节点
	visited := make(map[string]struct{})

	// BFS start here
	step := 1
	for len(beginSet) > 0 && len(endSet) > 0 { // fixme: 为什么是 && ？  beginSet 为空代表什么？所有节点都遍历完了？ beginSet endSet 其中一个为空就代表搜索完成了？ // 开始扩散
		if len(beginSet) > len(endSet) { // 为什么？ 从较小的集合扩散出去的节点比较少(从概率上讲)
			beginSet, endSet = endSet, beginSet
		}

		nextSet := make(map[string]struct{}) // 扩散后的点集合，即下一层的点集合
		for word := range beginSet {         // 扩散当前层所有节点

			visited[word] = struct{}{} // 添加到已访问过的节点集合

			for i := 0; i < len(word); i++ {
				for j := 'a'; j <= 'z'; j++ { // fixme: j <= 'z' 忘记 = 了。
					if byte(j) == word[i] { // 如果 nexWord == word
						continue
					}

					nexWord := word[:i] + string(j) + word[i+1:] // 扩散后的点，即下一层的点  // fixme: word[i+1:]

					if _, has := endSet[nexWord]; has {
						return step + 1
					}

					if _, has := wordSet[nexWord]; has { // 是否合法   扩散出来合法的节点
						if _, has = visited[nexWord]; !has { // 是否已经扩散(访问)过了？
							nextSet[nexWord] = struct{}{}
						}
					}
				}
			}
		}
		beginSet = nextSet
		step++
	}

	return 0
}
