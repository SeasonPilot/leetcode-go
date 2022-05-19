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
	al := len(a) - 1
	bl := len(b) - 1
	ai := al - bl - 1
	for i := bl; i >= 0; {
		if ai < 0 {
			a[al] = b[i]
			i--
		} else {
			if b[i] < a[ai] {
				a[al] = a[ai]
				ai--
			} else {
				a[al] = b[i]
				i--
			}
		}
		al--
	}
	return a
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
