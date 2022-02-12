// 选择排序 时间复杂度O(N^2)
package main

func ChooseSortAsc(array []int) []int {
	for i := 0; i <= len(array)-1; i++ {
		minPos := i
		for j := minPos + 1; j <= len(array)-1; j++ {
			if array[j] <= array[minPos] {
				minPos = j
			}
		}
		array[i], array[minPos] = array[minPos], array[i]
	}
	return array
}

func ChooseSortDesc(array []int) []int {
	for i := 0; i <= len(array)-1; i++ {
		minPos := i
		for j := minPos + 1; j <= len(array)-1; j++ {
			if array[j] >= array[minPos] {
				minPos = j
			}
		}
		array[i], array[minPos] = array[minPos], array[i]
	}
	return array
}

// func main() {
// 	ascResult := ChooseSortAsc([]int{55, 50, 18, 13, 16, 66, 34, 77, 23, 46})
// 	descResult := ChooseSortDesc([]int{55, 50, 18, 13, 16, 66, 34, 77, 23, 46})
// 	fmt.Println(ascResult)
// 	fmt.Println(descResult)
// }
