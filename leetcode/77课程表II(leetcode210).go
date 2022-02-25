package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := buildCoursesGraph(numCourses, prerequisites)
	visited, onPath, postOrder := make([]bool, numCourses), make([]bool, numCourses), []int{}
	var isCycle bool
	for i := 0; i < numCourses; i++ {
		findOrderTraverse(graph, i, visited, onPath, &isCycle, &postOrder)
	}
	if isCycle {
		return []int{}
	}

	arr := reverseArray(postOrder)
	return arr
}

func findOrderTraverse(graph [][]int, s int, visited, onPath []bool, isCycle *bool, postOrder *[]int) {
	if onPath[s] {
		*isCycle = true
	}
	if visited[s] || *isCycle {
		return
	}
	visited[s] = true
	onPath[s] = true
	for i := 0; i < len(graph[s]); i++ {
		findOrderTraverse(graph, graph[s][i], visited, onPath, isCycle, postOrder)
	}
	*postOrder = append(*postOrder, s)
	onPath[s] = false
}

func reverseArray(nums []int) []int {
	arr := make([]int, len(nums))
	for i, j := 0, len(nums)-1; i < len(arr) && j >= 0; {
		arr[i] = nums[j]
		i++
		j--
	}
	return arr
}
