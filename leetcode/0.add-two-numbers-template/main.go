package main

func main() {

}

//代码模板
//https://leetcode-cn.com/problems/add-to-array-form-of-integer/solution/jian-dan-yi-dong-javacpythonjs-pei-yang-a8ofe/
//【简单易懂Java/C++/Python/js/go】【培养算法思维】- 数组形式的整数加法
func addTwoNum(nums1 []int, nums2 []int) []int {
	res := make([]int, 0) // 结果
	carry := 0            // 进位
	l1, l2 := len(nums1)-1, len(nums2)-1
	for l1 >= 0 || l2 >= 0 { // 两个数组长度不同时，循环到较长的数组结束
		x, y := 0, 0 // 数组循环完时用 0 占位
		if l1 >= 0 {
			x = nums1[l1] // 如果数组未循环完则取对应位置的值
		}
		if l2 >= 0 {
			y = nums2[l2]
		}

		sum := x + y + carry
		res = append(res, sum%10)
		carry = sum / 10 // 进位

		l1--
		l2--
	}

	// 循环结束后，如果进位不为 0 则加入结果中
	if carry != 0 {
		res = append(res, carry)
	}

	// 反转 res 切片
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}

	return res
}
