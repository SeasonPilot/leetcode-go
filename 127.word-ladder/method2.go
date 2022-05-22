package main

// https://leetcode.com/problems/word-ladder/discuss/1042695/golang%3A-2-end-BFS-memory-99-3.9MB-speed-99-24ms-self-explanatory-comment
// 没看懂
func ladderLengthOptim(beginWord string, endWord string, wordList []string) int {
	endIdx := -1
	begIdx := -1
	for ii := range wordList {
		if wordList[ii] == endWord {
			endIdx = ii
		} else if wordList[ii] == beginWord {
			begIdx = ii
		}
	}

	if endIdx < 0 {
		return 0
	}

	// wordList is now all unvisited candidate node
	wordList = removeIndex(wordList, begIdx, endIdx)

	beginSet := []string{beginWord}
	endSet := []string{endWord}
	tmpSet := []string{} // visited???

	cnt := 1 // BFS starts here ???
	for len(beginSet) > 0 && len(endSet) > 0 {
		if len(beginSet) > len(endSet) {
			swapSet(&beginSet, &endSet)
		}

		tmpSet = tmpSet[:0]
		for _, word := range beginSet {

			// check if endSet and beginSet intersects
			for _, target := range endSet {
				if diffByOne(word, target) {
					return cnt + 1
				}
			}

			jj := len(wordList)
			for ii := 0; ii < jj; ii++ {
				target := wordList[ii]
				if diffByOne(word, target) {
					tmpSet = append(tmpSet, target)
					jj--
					wordList[ii] = wordList[jj]
					wordList[jj] = target
					// keep ii at same position
					ii--
				}
			}

			wordList = wordList[:jj]
		}

		// new frontline
		swapSet(&beginSet, &tmpSet)
		cnt++
	}

	return 0
}

func diffByOne(word1, word2 string) bool {
	diff := false
	for ii := range word1 {
		if word1[ii] != word2[ii] {
			if diff {
				return false
			}

			diff = true
		}
	}

	return diff
}

func swap(vec []string, ii, jj int) {
	tmp := vec[ii]
	vec[ii] = vec[jj]
	vec[jj] = tmp
}

// need to swapSet in place
func swapSet(vec1, vec2 *[]string) {
	tmp := *vec1
	*vec1 = *vec2
	*vec2 = tmp
}

// remove tail index first by swapping with the last node
func removeIndex(wordList []string, begIdx, endIdx int) []string {
	if endIdx < begIdx {
		tmp := begIdx
		begIdx = endIdx
		endIdx = tmp
	}

	if endIdx < 0 {
		return wordList
	}

	swap(wordList, endIdx, len(wordList)-1)
	wordList = wordList[:len(wordList)-1]

	if begIdx > 0 {
		swap(wordList, begIdx, len(wordList)-1)
		wordList = wordList[:len(wordList)-1]
	}

	return wordList
}
