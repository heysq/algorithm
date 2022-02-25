package leetcode

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	allPathsSourceTargetHelper(graph, []int{}, &res, 0)
	return res
}

func allPathsSourceTargetHelper(graph [][]int, onPath []int, res *[][]int, s int) {
	onPath = append(onPath, s)
	if s == len(graph)-1 {
		*res = append(*res, onPath)
		return
	}
	for _, i := range graph[s] {
		allPathsSourceTargetHelper(graph, onPath, res, i)
	}
	onPath = onPath[:len(onPath)-1]
}
