package b_huawei_leetcode

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				numIsLandsDfs(grid, i, j)
			}
		}
	}
	return res
}

func numIsLandsDfs(grid [][]byte, i, j int) {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	numIsLandsDfs(grid, i+1, j)
	numIsLandsDfs(grid, i-1, j)
	numIsLandsDfs(grid, i, j+1)
	numIsLandsDfs(grid, i, j-1)
}
