package leetcode

func numTrees(n int) int {

	var arr = make([][]int, n+1)
	for i := range arr {
		arr[i] = make([]int, n+1)
	}
	return numTreesCount(1, n, &arr)
}

func numTreesCount(low, high int, arr *([][]int)) int {
	if low > high {
		return 1
	}
	if val := (*arr)[low][high]; val != 0 {
		return val
	}
	var res int
	for i := low; i <= high; i++ {
		left := numTreesCount(low, i-1, arr)
		right := numTreesCount(i+1, high, arr)
		res += left * right
	}
	(*arr)[low][high] = res
	return res
}
