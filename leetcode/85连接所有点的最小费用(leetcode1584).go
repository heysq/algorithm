package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func minCostConnectPoints(points [][]int) int {
	connect := buildWeightArray(points)
	fmt.Println(connect.Connections)
	uf := NewUF(len(points))
	sort.Sort(connect)
	minCost := 0
	for _, val := range connect.Connections {
		from, to, weight := val[0], val[1], val[2]
		if uf.Connected(from, to) {
			continue
		}
		minCost += weight
		uf.Union(from, to)
	}
	return minCost
}

func buildWeightArray(points [][]int) *MyConnections {
	fmt.Println(points)
	weightArray := [][]int{}
	// 第几个点，具体点位
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			distance := math.Abs(float64(points[i][0]-points[j][0])) + math.Abs(float64(points[i][1]-points[j][1]))
			fmt.Println("i ", i, "j ", j, "dis ", distance)
			weightArray = append(weightArray, []int{i, j, int(distance)})
		}
	}
	return &MyConnections{Connections: weightArray}
}
