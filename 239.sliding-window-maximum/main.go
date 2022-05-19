package main

import "fmt"

func main() {
	nums := []int{100, 3, -1, -3, -5, -6, -7, -8}
	//nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow2(nums, k))
}

// 方法二：单调队列
func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

// 解法一 暴力解法 O(nk)
func maxSlidingWindow1(a []int, k int) []int {
	res := make([]int, 0, k) // 结果
	n := len(a)
	if n == 0 {
		return []int{}
	}
	// 从头开始滑动窗口
	for i := 0; i <= n-k; i++ {
		max := a[i]              // 初始值为第一个数值
		for j := 1; j < k; j++ { // 遍历当前窗口内的所有元素并比较大小
			if max < a[i+j] {
				max = a[i+j]
			}
		}
		res = append(res, max)
	}

	return res
}

// 解法二 双端队列 Deque
func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	window := make([]int, 0, k) // store the index of nums   最左边是 nums 中最大值的索引；单调队列
	result := make([]int, 0, len(nums)-k+1)

	for i, v := range nums { // if the left-most index is out of window, remove it
		// 窗口开始移动的时间和窗口最多能移动到哪
		if i >= k && window[0] <= i-k { // 遍历到窗口外时，开始移动窗口;窗口的左边界最大只能移动到 i-k
			window = window[1:] // 向右移动窗口
		}

		// 窗口内有值
		for len(window) > 0 && nums[window[len(window)-1]] < v { // for循环处理，窗口中存储的最后一个 index 对应的 nums < 当前值
			window = window[0 : len(window)-1] // 移出 Window 最后的元素
		}
		window = append(window, i) // 将当前 num 对应的 index 存储下来

		if i >= k-1 { // 第一次遍历到窗口中最后一个数时，开始返回结果
			result = append(result, nums[window[0]]) // the left-most is the index of max value in nums；返回最左边的元素
		}

		fmt.Println("窗口内的index: ", window)
	}
	return result
}
