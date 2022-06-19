package main

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("%#v\n", twoSum(nums, target))
}
