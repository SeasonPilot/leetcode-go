package main

// 双向 BFS
// 链接：https://leetcode.cn/problems/word-ladder/solution/dan-xiang-bfs-shuang-xiang-bfs-golangban-p1en/
func ladderLength3(beginWord string, endWord string, wordList []string) int {
	// setep1：先将 wordList 放到哈希表里，便于判断某个单词是否在 wordList 里
	wordSet := make(map[string]struct{})
	for _, wd := range wordList {
		wordSet[wd] = struct{}{}
	}
	// 如果 endWord 不存在于 wordList 中，则无法转换
	if _, has := wordSet[endWord]; !has {
		return 0
	}

	// step2：已经访问过的 word 添加到 visited 哈希表里
	visited := make(map[string]struct{})
	// 分别用左边和右边扩散的哈希表代替单向 BFS 里的队列，它们在双向 BFS 的过程中交替使用
	beginSet, endSet := make(map[string]struct{}), make(map[string]struct{})
	beginSet[beginWord] = struct{}{}
	endSet[endWord] = struct{}{}

	// step3：执行双向 BFS，左右交替扩散的步数之和为所求
	step := 1 // 相当于 Java 版本里面的 len
	for len(beginSet) > 0 && len(endSet) > 0 {
		// 优先选择小的哈希表进行扩散，考虑到的情况更少，此处保证 beginVisited 是较小的集合
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}

		// 逻辑到这里，保证 beginVisited 是相对较小的集合，nextVisited 在扩散完成以后，会成为新的 beginVisited
		nextSet := make(map[string]struct{}) // 相当于 Java 版本里面的 temp
		for word := range beginSet {
			if changeWordByOneLetter(word, endSet, visited, wordSet, nextSet) {
				return step + 1
			}
		}
		// 原来的 beginSet 废弃，从 nextSet 开始新的双向 BFS
		beginSet = nextSet
		step++
	}
	return 0
}

// changeWordByOneLetter4 尝试对 word 修改每一个字符，看看是不是能落在 endSet 中，扩展得到的新的 word 添加到 nextSet 里
func changeWordByOneLetter(word string, endSet, visited, wordSet, nextSet map[string]struct{}) bool {
	for i := 0; i < len(word); i++ {
		for k := 'a'; k <= 'z'; k++ {
			if byte(k) == word[i] {
				continue
			}
			nextWord := word[:i] + string(k) + word[i+1:]

			if _, has := wordSet[nextWord]; has {
				if _, has := endSet[nextWord]; has {
					return true
				}
				if _, has := visited[nextWord]; !has {
					nextSet[nextWord] = struct{}{}
					visited[nextWord] = struct{}{}
				}
			}
		}
	}

	return false
}
