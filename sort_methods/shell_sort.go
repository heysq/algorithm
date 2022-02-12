package main

func ShellSortAsc(array []int) []int {
	for gap := getArrayGap(array); gap > 0; gap = (gap - 1) / 3 {
		for i := gap; i < len(array); i++ {
			for j := i; (j - gap) >= 0; j -= gap {
				if array[j] < array[j-gap] {
					array[j], array[j-gap] = array[j-gap], array[j]
				}
			}
		}
	}
	return array
}

func ShellSortDesc(array []int) []int {
	for gap := getArrayGap(array); gap > 0; gap = (gap - 1) / 3 {
		// 从第一个gap开始，一直到数组结尾。
		for i := gap; i < len(array); i++ {
			for j := i; (j - gap) >= 0; j -= gap {
				if array[j] > array[j-gap] {
					array[j], array[j-gap] = array[j-gap], array[j]
				}
			}
		}
	}
	return array
}

func getArrayGap(array []int) int {
	var gap = 1
	for gap <= len(array)/3 {
		gap = 3*gap + 1
	}
	return gap
}

// func main() {
// 	array := helper.RandomIntArray(10, 1000)
// 	fmt.Println(array)
// 	result := ShellSortAsc(array)
// 	fmt.Println(result)
// 	for index, _ := range array {

// 	}

// }
