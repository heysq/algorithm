package leetcode

func solve(board [][]byte) {
	m, n := len(board), len(board[0]) // 行数 列数
	for i := 0; i < m; i++ {
		solveDFS(board, i, 0)
		solveDFS(board, i, n-1)
	}
	for j := 0; j < n; j++ {
		solveDFS(board, 0, j)
		solveDFS(board, m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func solveDFS(board [][]byte, i, j int) {
	m, n := len(board), len(board[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}

	if board[i][j] != 'O' {
		return
	}

	board[i][j] = '#'

	solveDFS(board, i-1, j)
	solveDFS(board, i+1, j)
	solveDFS(board, i, j-1)
	solveDFS(board, i, j+1)
}
