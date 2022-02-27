package leetcode

func validTree(n int, edges [][]int) bool {
	uf := NewUF(n)
	for _, i := range edges {
		from := i[0]
		to := i[1]
		if uf.Connected(from, to) {
			return false
		}
		uf.Union(from, to)
	}
	return uf.Count() == 1
}
