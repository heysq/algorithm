package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	var isCycle bool
	graph := buildCoursesGraph(numCourses, prerequisites)
	for i := 0; i < numCourses; i++ {
		canFinishTraverse(graph, i, visited, onPath, &isCycle)
	}
	return isCycle
}

func canFinishTraverse(graph [][]int, s int, visited, onPath []bool, isCycle *bool) {
	if onPath[s] {
		*isCycle = true
	}
	if visited[s] || *isCycle {
		return
	}
	visited[s] = true
	onPath[s] = true
	for _, i := range graph[s] {
		canFinishTraverse(graph, i, visited, onPath, isCycle)
	}
	onPath[s] = false
}

func buildCoursesGraph(n int, prerequisites [][]int) [][]int {
	arr := make([][]int, n)
	for i := 0; i < len(prerequisites); i++ {
		p := prerequisites[i][0]
		q := prerequisites[i][1]
		arr[p] = append(arr[p], q)
	}
	return arr
}
