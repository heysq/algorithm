package leetcode

import (
	"strings"
)

func solveNQueens(n int) [][]string {
	var result [][]string
	var grid = make([][]string, n)
	for i := range grid {
		grid[i] = make([]string, n)
		for j := 0; j < n; j++ {
			grid[i][j] = "."
		}
	}
	// fmt.Println("grid --->", grid)
	solveNQueensHelper(grid, &result, 0)
	return result
}

func solveNQueensHelper(grid [][]string, result *[][]string, row int) {
	if row == len(grid) {
		arr := make([]string, row)
		for i := range arr {
			arr[i] = strings.Join(grid[i], "")
		}
		// fmt.Println("arr -->", arr)
		*result = append(*result, arr)
		return
	}
	n := len(grid[row])
	for col := 0; col < n; col++ {
		if !isSolveNQueueValid(grid, row, col) {
			continue
		}
		grid[row][col] = "Q"
		solveNQueensHelper(grid, result, row+1)
		grid[row][col] = "."
	}
}

// 同一行 不能有皇后
// 左上不能有皇后
// 右上不能有皇后
func isSolveNQueueValid(grid [][]string, row, col int) bool {
	// 同一行不能有皇后
	for i := 0; i < len(grid); i++ {
		if grid[i][col] == "Q" {
			return false
		}
	}

	// 左上不能有皇后
	for r, c := row-1, col-1; r >= 0 && c >= 0; {
		if grid[r][c] == "Q" {
			return false
		}
		r--
		c--
	}

	// 右上不能有皇后
	for r, c := row-1, col+1; r >= 0 && c < len(grid); {
		if grid[r][c] == "Q" {
			return false
		}
		r--
		c++
	}
	return true
}
