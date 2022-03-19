package leetcode

// func nthUglyNumber(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	h := &SmallHeap{arr: []int{1}, size: 1}
// 	adjustSmallHeap(h, 0)
// 	count := 0
// 	hash := make(map[int]struct{})
// 	for h.size != 0  {
// 		cut := h.Poll()
// 		count++
// 		if count == n {
// 			return cut
// 		}
// 		for _, val := []int{2,3,5}{
// 			if _, ok := hash[cut * val]; !ok{
// 				h.Add(cut *val)
// 				hash[cut*val] = struct{}{}
// 			}
// 		}
// 	}
// 	return -1
// }

// type SmallHeap struct {
// 	arr  []int
// 	size int
// }

// func adjustSmallHeap(heap *SmallHeap, parent int) {
// 	left := parent*2 + 1
// 	right := parent*2 + 2

// 	minNode := parent

// 	if left < heap.size && heap.arr[minNode] > heap.arr[left] {
// 		minNode = left
// 	}

// 	if right < heap.size && heap.arr[minNode] > heap.arr[right] {
// 		minNode = right
// 	}

// 	if minNode != parent {
// 		heap.arr[minNode], heap.arr[parent] = heap.arr[parent], heap.arr[minNode]
// 		adjustSmallHeap(heap, minNode)
// 	}
// }

// func (h *SmallHeap) Poll() int {
// 	x := h.arr[0]
// 	h.arr = h.arr[1:]
// 	h.size--
// 	adjustSmallHeap(h, 0)
// 	return x
// }

// // 小根堆插入
// func (h *SmallHeap) Add(n int) {
// 	h.arr = append(h.arr, n)
// 	h.size++
// 	if len(h.arr) == 1 {
// 		return
// 	}
// 	adjustSmallHeapUp(h, h.size-1)
// }

// func adjustSmallHeapUp(heap *SmallHeap, n int) {
// 	parent := (n - 1) / 2
// 	if heap.arr[parent] > heap.arr[n] {
// 		heap.arr[n], heap.arr[parent] = heap.arr[parent], heap.arr[n]
// 		adjustSmallHeapUp(heap, parent)

// 	}
// }
