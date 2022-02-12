// 选择排序 时间复杂度O(N^2)
package main

import (
	"fmt"
)

//func BubbleSortAsc(array []int) []int {
//	for i := 0; i < len(array)-1; i++ {
//		for j := i + 1; j <= len(array)-1; j++ {
//			if array[i] > array[j] {
//				array[i], array[j] = array[j], array[i]
//			}
//		}
//	}
//	return array
//}

func BubbleSortAsc(array []int) []int {
	for i := len(array) - 1; i > 0; i-- {
		isSorted := true
		fmt.Println(array)
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
	return array
}

func BubbleSortDesc(array []int) []int {
	for i := len(array) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}

// func main() {
// 	var array []int
// 	array = helper.RandomIntArray(10, 1000)
// 	fmt.Println(array)
// 	ascResult := BubbleSortAsc([]int{1, 2, 3, 4, 5}) // 最好情况下是O(n)的，可以遍历发现是有序的，可以取消掉后边无效的遍历
// 	fmt.Println(ascResult)

// 	array = helper.RandomIntArray(10, 100)
// 	fmt.Println(array)
// 	descResult := BubbleSortDesc(array)
// 	fmt.Println(descResult)
// }
