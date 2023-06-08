package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	a := []int{2, 5, 7, 9, 0, 0, 0}
	b := []int{1, 3, 6}
	fmt.Println("hello https://tool.lu/", Asc(a, b))
}

/*
	 a = [ 2, 5, 7, 9, 0, 0, 0],
	 b = [ 1, 3, 6],
	说明：
	1、a 中 0 是无效数据，0的长度与 b 的长度一致，0在最后
	2、a 和 b 都是有序的，从小到大
	要求：以最小的时间和空间代价合并如下两个数组，并且保证有序，从小到大
*/
func Asc(a, b []int) []int {
	al := len(a) - 1 // 获取两个切片的长度，并分别存储在al和bl变量中。
	bl := len(b) - 1
	ai := al - bl - 1      // 定义一个ai变量，它表示第一个切片中用于比较的元素的索引（index），初始值为al-bl-1。
	for i := bl; i >= 0; { // 从后 往前遍历第二个切片中的每个元素，用i变量表示当前元素的索引。
		if ai < 0 { // 对于每个元素，如果ai小于0，说明第一个切片中已经没有元素可以比较了，直接将第二个切片中的当前元素赋值给第一个切片中的最后一个位置（由al表示），然后i减一。
			a[al] = b[i]
			i--
		} else { // 否则，如果ai不小于0，说明第一个切片中还有元素可以比较，那么就比较第二个切片中的当前元素和第一个切片中由ai表示的元素的大小。
			if b[i] < a[ai] { // 如果第二个切片中的当前元素小于第一个切片中由ai表示的元素，那么就将第一个切片中由ai表示的元素赋值给第一个切片中的最后一个位置（由al表示），然后ai减一。
				a[al] = a[ai]
				ai--
			} else { // 否则，如果第二个切片中的当前元素大于或等于第一个切片中由ai表示的元素，那么就将第二个切片中的当前元素赋值给第一个切片中的最后一个位置（由al表示），然后i减一。
				a[al] = b[i]
				i--
			}
		}
		al-- // 最后，无论哪种情况，都要将al减一，以便下一次循环时更新第一个切片中的最后一个位置。
	}
	return a // 当遍历完第二个切片中的所有元素后，函数返回第一个切片作为结果。
}

//golang删除文件中空行
func DeleteBlankFile(srcFilePah string, destFilePath string) error {
	srcFile, err := os.OpenFile(srcFilePah, os.O_RDONLY, 0666)
	defer srcFile.Close()
	if err != nil {
		return err
	}
	srcReader := bufio.NewReader(srcFile)
	destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer destFile.Close()
	if err != nil {
		return err
	}
	for {
		str, err := srcReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print("The file end is touched.")
				break
			} else {
				return err
			}
		}
		if strings.HasSuffix(str, " \r\n") {
			continue
		}
		if strings.HasPrefix(str, "\r\n") {
			continue
		}
		fmt.Println(len(str))
		fmt.Print(str)
		destFile.WriteString(str)
	}
	return nil
}
