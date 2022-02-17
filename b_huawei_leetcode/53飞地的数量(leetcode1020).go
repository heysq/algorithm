package b_huawei_leetcode

func numEnclaves(grid [][]int) int {
	m, n, res := len(grid), len(grid[0]), 0

	for i := 0; i < m; i++ {
		numEnclavesDFS(grid, i, n-1)
		numEnclavesDFS(grid, i, 0)
	}

	for j := 0; j < n; j++ {
		numEnclavesDFS(grid, 0, j)
		numEnclavesDFS(grid, m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func numEnclavesDFS(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	numEnclavesDFS(grid, i+1, j)
	numEnclavesDFS(grid, i-1, j)
	numEnclavesDFS(grid, i, j+1)
	numEnclavesDFS(grid, i, j-1)
}
