package main

// import (
// 	"fmt"
// 	"leet_code_go/helper"
// )

func QuickSort(array []int, left, right int) []int {
	if left >= right {
		return array
	}
	pivot := partition(array, left, right)
	QuickSort(array, left, pivot-1)
	QuickSort(array, pivot+1, right)
	return array
}

func partition(array []int, left, right int) int {
	pivot := array[right]
	var l, r = left, right - 1
	for l <= r {
		for l <= r && array[l] <= pivot {
			l++
		}
		for l <= r && array[r] > pivot {
			r--
		}

		if l < r {
			array[l], array[r] = array[r], array[l]
		}
	}
	array[l], array[right] = array[right], array[l]
	return l
}
// func main() {
// 	array := helper.RandomIntArray(10000, 100000)
// 	cArray := helper.CopyArray(array)
// 	result := QuickSort(cArray, 0, len(cArray)-1)
// 	fmt.Println(helper.ArraySortCompare(result, array))
// }
