package main

import (
	"fmt"
	"testing"
)

func TestRemoveInterval(t *testing.T) {
	intervals := [][]int{{0, 2}, {3, 4}, {5, 7}}
	toBeRemoved := []int{1, 6}
	//[[0, 1], [6, 7]]
	fmt.Println(removeInterval(intervals, toBeRemoved))
	fmt.Println(removeInterval1(intervals, toBeRemoved))
}
