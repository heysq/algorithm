package leetcode

func maxAreaOfIsland(grid [][]int) int {
	m, n, res := len(grid), len(grid[0]), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := maxAreaOfIslandDFS(grid, i, j)
				if area > res {
					res = area
				}
			}
		}
	}
	return res
}

func maxAreaOfIslandDFS(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return maxAreaOfIslandDFS(grid, i-1, j) +
		maxAreaOfIslandDFS(grid, i+1, j) +
		maxAreaOfIslandDFS(grid, i, j-1) +
		maxAreaOfIslandDFS(grid, i, j+1) + 1
}
