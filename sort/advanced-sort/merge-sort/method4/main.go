package main

import "fmt"

// 对齐小魔王 PPT 上的代码
// 对当前的序列（start到end）进行归并排序
func mergeSort(arr []int, start, end int) {
	if end <= start { // 递归的出口：不能再二分了，返回
		return
	}
	//mid := start + (end-start)>>1 // 当前序列的中点索引
	mid := (start + end) >> 1 // 当前序列的中点索引

	mergeSort(arr, start, mid) // 递归左序列
	mergeSort(arr, mid+1, end) // 递归右序列

	i := start
	j := mid + 1 // 复原 i j 指针，因为现在要合并左右序列

	temp := make([]int, end-start+1) // 辅助数组，存放合并排序的数

	k := 0                     // 从0开始
	for i <= mid && j <= end { // 如果 i j 都没越界
		if arr[i] <= arr[j] { // arr[i]更小
			temp[k] = arr[i] // 取arr[i]，确定了temp[k]
			k++              // 确定下一个
			i++              // 考察下一个i，j不动
		} else {
			temp[k] = arr[j]
			k++
			j++
		}
	}
	for i <= mid { // 如果 i 没越界，j越界了
		temp[k] = arr[i] // i 和 i右边的都取过来
		k++              // 确定下一个数
		i++
	}
	for j <= end { // j 没越界，i越界了
		temp[k] = arr[j] // j 和 j右边的都取过来
		k++              // 确定下一个数
		j++
	}

	for p := 0; p < len(temp); p++ {
		arr[start+p] = temp[p]
	}
}

//链接：https://leetcode.cn/problems/reverse-pairs/solution/shou-hua-tu-jie-yi-bu-yi-bu-jie-xi-gui-bing-pai-xu/

func main() {
	s := []int{6, 7, 8, 10, 4, 6, 99}
	mergeSort(s, 0, len(s)-1)
	fmt.Println(s)
}
