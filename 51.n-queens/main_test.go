package main

import (
	"fmt"
	"testing"
)

func TestQueens(t *testing.T) {
	//fmt.Println(solveNQueens(8))
	s := solveNQueens(8)
	fmt.Println(len(s))
}
