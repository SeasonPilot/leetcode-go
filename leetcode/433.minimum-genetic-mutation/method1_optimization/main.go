package method1

// 最终使用此方法
// BFS
func minMutation(start string, end string, bank []string) int {
	// bank slice 转 hashmap
	bankSet := make(map[string]struct{}, len(bank))
	for _, b := range bank {
		bankSet[b] = struct{}{}
	}

	// 特殊条件判断
	if start == end { // fixme: 没有判断边界条件
		return 0
	}
	if _, has := bankSet[end]; !has {
		return -1
	}

	// queue
	q := []string{start}
	// visited

	step := 1
	for len(q) != 0 {
		// pop
		var nextQ []string

		for _, cur := range q {
			for i, x := range cur {
				for _, y := range "ACGT" {
					if x != y {
						ns := cur[:i] + string(y) + cur[i+1:]

						if _, has := bankSet[ns]; has { // if ns in bank    将 bank 转换成 hashmap 方便快速查找 ns 是否合法
							if ns == end { // 出口
								return step
							}
							delete(bankSet, ns)
							nextQ = append(nextQ, ns)
						}
					}
				}
			}
		}

		q = nextQ
		step++ // 扩散完一层后 step + 1
	}
	return -1 // fixme: 如果无法完成此基因变化，返回 -1
}
