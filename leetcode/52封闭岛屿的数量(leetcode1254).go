package leetcode

func closedIsland(grid [][]int) int {

	m, n, res := len(grid), len(grid[0]), 0

	for i := 0; i < m; i++ {
		closedIslandDFS(grid, i, 0)
		closedIslandDFS(grid, i, n-1)
	}

	for j := 0; j < n; j++ {
		closedIslandDFS(grid, 0, j)
		closedIslandDFS(grid, m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				closedIslandDFS(grid, i, j)
			}
		}
	}
	return res
}

func closedIslandDFS(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	closedIslandDFS(grid, i+1, j)
	closedIslandDFS(grid, i-1, j)
	closedIslandDFS(grid, i, j-1)
	closedIslandDFS(grid, i, j+1)
}
