package leetcode

import (
	"fmt"
	"testing"
)

func Test_solve(t *testing.T) {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Println(board)
}
