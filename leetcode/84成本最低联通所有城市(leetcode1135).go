package leetcode

import (
	"sort"
)

type MyConnections struct {
	Connections [][]int
}

func (m *MyConnections) Less(i, j int) bool {
	return m.Connections[i][2] < m.Connections[j][2]
}

func (m *MyConnections) Len() int {
	return len(m.Connections)
}

func (m *MyConnections) Swap(i, j int) {
	m.Connections[i], m.Connections[j] = m.Connections[j], m.Connections[i]
}

func miniumCost(n int, connections [][]int) int {
	uf := NewUF(n + 1)
	connect := MyConnections{Connections: connections}
	sort.Sort(&connect)
	var minCost int
	for _, c := range connect.Connections {
		from, to, weight := c[0], c[1], c[2]
		if uf.Connected(from, to) {
			continue
		}
		minCost += weight
		uf.Union(from, to)
	}
	if uf.Count() == 2 {
		return minCost
	}
	return -1
}
