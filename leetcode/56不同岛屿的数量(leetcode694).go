package leetcode

func numDistinctIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var data = make(map[string]struct{}, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				var res string
				numDistinctIslandsDFS(grid, i, j, &res, "666")
				data[res] = struct{}{}
			}
		}
	}
	return len(data)
}

func numDistinctIslandsDFS(grid [][]int, i, j int, str *string, dir string) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	*str = *str + dir + ","
	numDistinctIslandsDFS(grid, i-1, j, str, "1")
	numDistinctIslandsDFS(grid, i+1, j, str, "2")
	numDistinctIslandsDFS(grid, i, j-1, str, "3")
	numDistinctIslandsDFS(grid, i, j+1, str, "4")
	*str = *str + "-" + dir + ","
}
