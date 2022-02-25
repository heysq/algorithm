package leetcode

import "fmt"

func possibleBipartition(n int, dislikes [][]int) bool {
	graph := possibleBipartitionBuildGraph(n, dislikes)
	fmt.Println(graph)
	// return false

	visited, color := make([]bool, n+1), make([]bool, n+1)
	var ok = true
	for i := 1; i < len(graph); i++ {
		if !visited[i] {
			possibleBipartDFS(graph, i, visited, color, &ok)
		}
	}
	return ok
}

func possibleBipartDFS(graph [][]int, s int, visited, color []bool, ok *bool) {
	if !*ok {
		return
	}

	visited[s] = true
	for _, i := range graph[s] {
		if !visited[i] {
			visited[i] = true
			color[i] = !color[s]
			possibleBipartDFS(graph, i, visited, color, ok)
		} else if color[i] == color[s] {
			*ok = false
		}
	}
}

func possibleBipartitionBuildGraph(n int, dislikes [][]int) [][]int {
	var graph = make([][]int, n+1)
	for _, i := range dislikes {
		from := i[0]
		to := i[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	return graph
}
