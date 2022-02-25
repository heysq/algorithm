package leetcode

type UF struct {
	count int
	arr   []int
	size  []int
}

func NewUF(n int) *UF {
	uf := UF{
		arr:  make([]int, n),
		size: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.arr[i] = i
		uf.size[i] = i
		uf.count++
	}
	return &uf
}

func (uf *UF) Find(n int) int {
	x := uf.arr[n]
	for x != n {
		x = uf.arr[uf.arr[x]] // 每次压缩一半
	}
	return x
}

func (uf *UF) Union(n, m int) {
	rootN := uf.Find(n)
	rootM := uf.Find(m)

	if rootN == rootM {
		return
	}

	if uf.size[rootN] > uf.size[rootM] {
		uf.arr[rootM] = rootN
		uf.size[rootN] += uf.size[rootM]
	} else {
		uf.arr[rootN] = rootM
		uf.size[rootN] += uf.size[rootM]
	}
	uf.count--
}

func (uf *UF) Connected(n, m int) bool {
	return uf.Find(n) == uf.Find(m)
}
