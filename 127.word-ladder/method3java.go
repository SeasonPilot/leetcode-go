package main

// 链接：https://leetcode.cn/problems/word-ladder/solution/dan-xiang-bfs-shuang-xiang-bfs-golangban-p1en/
// 双向 BFS   对齐覃超的 Java 版本思路。 practice 使用的是这个版本。
func ladderLength4(beginWord string, endWord string, wordList []string) int {
	// step1：先将 wordList 放到哈希表里，便于判断某个单词是否在 wordList 里
	wordSet := make(map[string]struct{}) // 使用哈希map 便于查单词是否存在于wordlist，时间复杂度为O(1)
	for _, wd := range wordList {
		wordSet[wd] = struct{}{}
	}
	// 如果 endWord 不存在于 wordList 中，则无法转换
	if _, has := wordSet[endWord]; !has {
		return 0
	}

	// step2：已经访问过的 word 添加到 visited 哈希表里
	visited := make(map[string]struct{}) // ????  BFS 节点是否已经被访问过

	// 分别用左边和右边扩散的哈希表代替单向 BFS 里的队列，它们在双向 BFS 的过程中交替使用
	beginSet, endSet := make(map[string]struct{}), make(map[string]struct{})
	beginSet[beginWord] = struct{}{} // 从 begin word 扩散出来的 Set
	endSet[endWord] = struct{}{}

	// step3：执行双向 BFS，左右交替扩散的步数之和为所求
	step := 1                                  // 相当于 Java 版本里面的 len。    水波纹扩散的一层？？
	for len(beginSet) > 0 && len(endSet) > 0 { // 相当于单向 BFS 的 queue
		// 优先选择小的哈希表进行扩散，考虑到的情况更少，此处保证 beginVisited 是较小的集合
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}

		// 逻辑到这里，保证 beginVisited 是相对较小的集合，nextVisited 在扩散完成以后，会成为新的 beginVisited
		nextSet := make(map[string]struct{}) // 相当于 Java 版本里面的 temp
		for word := range beginSet {         // 扩散 beginSet

			for i := 0; i < len(word); i++ {
				for k := 'a'; k <= 'z'; k++ {
					if byte(k) == word[i] {
						continue
					}

					target := word[:i] + string(k) + word[i+1:] // fixme: ait

					if _, has := endSet[target]; has { // beginSet endSet 相交了
						return step + 1
					}
					if _, has := wordSet[target]; has { // fixme: ait 是否在 wordSet 中。hot 在 wordSet 中 进入语句内
						if _, has = visited[target]; !has {
							nextSet[target] = struct{}{} // 下一次要扩散的 set
							visited[target] = struct{}{}
						}
					}
				}
			}
		}

		// 原来的 beginSet 废弃，从 nextSet 开始新的双向 BFS
		beginSet = nextSet
		step++
	}

	return 0
}
