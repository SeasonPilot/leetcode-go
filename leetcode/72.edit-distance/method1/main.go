package main

//链接：https: //leetcode.cn/problems/edit-distance/solution/dong-tai-gui-hua-gun-dong-shu-zu-by-acra-htab/
func minDistance(word1 string, word2 string) int {
	if len(word2)*len(word1) == 0 {
		return len(word1) + len(word2)
	}

	first, second := make([]int, len(word2)+1), make([]int, len(word2)+1)
	for i := range first {
		first[i] = i
	}
	for i := range word1 {
		second[0] = i + 1
		for j := range word2 {
			b := j + 1
			if word1[i] == word2[j] {
				second[b] = first[b-1]
			} else {
				second[b] = min(first[b]+1, min(first[b-1]+1, second[b-1]+1))
			}
		}
		copy(first, second)
	}
	return second[len(word2)]
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
