package method1

/*
由 cody 根据 https://labuladong.github.io/algo/di-ling-zh-bfe1b/hui-su-sua-c26da/#:~:text=javascript%20%F0%9F%A4%96-,class%20Solution%20%7B,%7D,-%F0%9F%91%BE%20%E4%BB%A3%E7%A0%81%E5%8F%AF%E8%A7%86
翻译
*/

//回溯算法
func permute(nums []int) [][]int {
	var res [][]int

	// 记录「路径」
	var track []int // 「路径」中的元素会被标记为 true,避免重复使用
	used := make([]bool, len(nums))

	backtrack(nums, track, used, &res)
	return res
}

// 路径:记录在 track 中
// 选择列表:nums 中不存在于 track 的那些元素(used[i] 为 false)
// 结束条件:nums 中的元素全都在 track 中出现
func backtrack(nums []int, track []int, used []bool, res *[][]int) { //fixme: res 要是用指针类型
	// 触发结束条件
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 排除不合法的选择
		if used[i] {
			// nums[i] 已经在 track 中,跳过
			continue
		}
		// 做选择
		track = append(track, nums[i])
		used[i] = true
		// 进入下一层决策树
		backtrack(nums, track, used, res)
		// 取消选择
		track = track[:len(track)-1]
		used[i] = false
	}
}

/*
为什么加入解集时，要将数组内容拷贝到一个新的数组里，再加入解集？

因为该 path 变量存的是地址引用，结束当前递归时，将它加入 res 后，该算法还要进入别的递归分支继续搜索，还要继续将这个 path 传给别的递归调用，
它所指向的内存空间还要继续被操作，所以 res 中的 path 的内容会被改变，这就不对。
所以要弄一份当前的拷贝，放入 res，这样后续对 path 的操作，就不会影响已经放入 res 的内容。

作者：xiao_ben_zhu
链接：https://leetcode.cn/problems/permutations/solution/chou-xiang-cheng-jue-ce-shu-yi-ge-pai-lie-jiu-xian/
*/

/*
因为切片是引用类型，原答案里递归调用时参数直接传的是path，因此不同层的递归调用栈里的path指向相同的底层数组，这就导致修改slice时，
slice底层的数组修改，如果不copy，保存一份快照的话，result保存的path可能会指向同一个数组，即已经保存的path会因后续的修改而改变。
这里不用copy是因为，没有修改path，即没有用append返回的新slice替换path，所以每个调用栈帧里的path都是一个新的slice。

https://leetcode.cn/problems/permutations/solution/chou-xiang-cheng-jue-ce-shu-yi-ge-pai-lie-jiu-xian/674643
*/
