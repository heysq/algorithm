package leetcode

func isBipartite(graph [][]int) bool {
	visited := make([]bool, len(graph))
	color := make([]bool, len(graph))
	var ok = true
	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			isBipartiteBFS(graph, i, visited, color, &ok)
		}
	}
	return ok
}

func isBipartiteDFS(graph [][]int, s int, visited, color []bool, ok *bool) {
	if visited[s] {
		return
	}
	visited[s] = true
	for _, v := range graph[s] {
		if !visited[v] {
			color[v] = !color[s]
			isBipartiteDFS(graph, v, visited, color, ok)
		} else if color[s] == color[v] {
			*ok = false
		}
	}
}

func isBipartiteBFS(graph [][]int, s int, visited, color []bool, ok *bool) {
	if visited[s] {
		return
	}
	arr := []int{}
	arr = append(arr, s)
	visited[s] = true
	for len(arr) != 0 && *ok {
		v := arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		for _, i := range graph[v] {
			if !visited[i] {
				visited[i] = true
				arr = append(arr, i)
				color[i] = !color[v]
			} else if color[v] == color[i] {
				*ok = false
			}
		}
	}
}
