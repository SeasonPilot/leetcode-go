package practice1

//回溯算法
func permute(nums []int) [][]int {
	var result [][]int
	var track []int
	//var used []bool //fixme:要初始化切片
	used := make([]bool, len(nums))

	backtrack(nums, track, used, &result)
	return result
}

func backtrack(nums []int, track []int, used []bool, res *[][]int) { //fixme:res要用指针。 - 函数传参slice传的是个struct，函数内部对slice进行append，因为外部struct的len不变，所以外部感知不到内部的append。
	if len(track) == len(nums) {
		//var temp []int
		temp := make([]int, len(track)) //fixme:要初始化切片,copy() 的 dst 不能为 nil
		copy(temp, track)               //destSlice可以为nil，但是这样的话，copy()函数就不会有任何效果，因为nil切片没有底层数组可以复制元素。
		*res = append(*res, temp)
		return
	}

	for i, n := range nums {
		if used[i] == true { //fixme:
			continue
		}

		track = append(track, n)
		used[i] = true
		backtrack(nums, track, used, res)
		track = track[:len(track)-1]
		used[i] = false
	}
}
