package b_huawei_leetcode

func findCircleNum(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	for i, v := range vis {
		if !v {
			ans++
			findCircleNumDFS(isConnected, i, vis)
		}
	}
	return
}

func findCircleNumDFS(isConnected [][]int, from int, vis []bool) {
	vis[from] = true
	for to, conn := range isConnected[from] {
		if conn == 1 && !vis[to] {
			findCircleNumDFS(isConnected, to, vis)
		}
	}
}
