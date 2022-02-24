package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	//1.去除非数字字母
	filterS := filterNonNumberAndChar(s)

	//2.字母全部转换为小写
	lowerS := strings.ToLower(filterS)

	//3.反转字符串进行对比
	return Reverse(lowerS) == lowerS
}

func filterNonNumberAndChar(s string) string {
	reg, err := regexp.Compile("[^\\w]")
	//reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return ""
	}
	return reg.ReplaceAllString(s, "")
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return string(a)
}

func main() {
	//example := "#tesGSGt252!$!"
	example := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(example))
}
