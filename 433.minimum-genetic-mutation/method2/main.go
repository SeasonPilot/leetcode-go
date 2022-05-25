package main

//链接：https://leetcode.cn/problems/minimum-genetic-mutation/solution/zui-xiao-ji-yin-bian-hua-by-leetcode-sol-lhwy/
func minMutation(start, end string, bank []string) int {
	if start == end {
		return 0
	}
	bankSet := map[string]struct{}{}
	for _, s := range bank {
		bankSet[s] = struct{}{}
	}
	if _, ok := bankSet[end]; !ok {
		return -1
	}

	q := []string{start}
	for step := 0; q != nil; step++ {
		tmp := q
		q = nil
		for _, cur := range tmp {
			for i, x := range cur {
				for _, y := range "ACGT" {
					if y != x {
						nxt := cur[:i] + string(y) + cur[i+1:]
						if _, ok := bankSet[nxt]; ok {
							if nxt == end {
								return step + 1
							}
							delete(bankSet, nxt)
							q = append(q, nxt)
						}
					}
				}
			}
		}
	}
	return -1
}
