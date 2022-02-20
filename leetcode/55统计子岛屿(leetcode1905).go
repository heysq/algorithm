package leetcode

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n, res := len(grid1), len(grid1[0]), 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				countSubIslandsDFS(grid2, i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				countSubIslandsDFS(grid2, i, j)
			}
		}
	}
	return res
}

func countSubIslandsDFS(grid2 [][]int, i, j int) {
	m, n := len(grid2), len(grid2[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}

	if grid2[i][j] == 0 {
		return

	}
	grid2[i][j] = 0
	countSubIslandsDFS(grid2, i-1, j)
	countSubIslandsDFS(grid2, i+1, j)
	countSubIslandsDFS(grid2, i, j-1)
	countSubIslandsDFS(grid2, i, j+1)
}
