package leetcode

/*
构建大根堆
*/
func smallestK(arr []int, k int) []int {

	if len(arr) == 0 || k < 0 {
		return []int{}
	}

	mhp := &BigHeap{arr: arr[:k], size: k}
	for i := mhp.size / 2; i >= 0; i-- {
		adjustBigHeap(mhp, i)
	}

	for _, val := range arr[k:] {
		if val < mhp.arr[0] {
			mhp.arr[0] = val
			adjustBigHeap(mhp, 0)
		}
	}
	return mhp.arr
}

type BigHeap struct {
	arr  []int
	size int
}

func adjustBigHeap(heap *BigHeap, parent int) {
	left := parent*2 + 1
	right := parent*2 + 2
	maxNode := parent

	if left < heap.size && heap.arr[maxNode] < heap.arr[left] {
		maxNode = left
	}

	if right < heap.size && heap.arr[maxNode] < heap.arr[right] {
		maxNode = right
	}

	if maxNode != parent {
		heap.arr[maxNode], heap.arr[parent] = heap.arr[parent], heap.arr[maxNode]
		adjustBigHeap(heap, maxNode)
	}
}
