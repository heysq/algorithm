package main

func InsertSortAsc(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i - 1; j >= 0; j-- {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func InsertSortDesc(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i - 1; j >= 0; j-- {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

// func main() {
// 	array := helper.RandomIntArray(10, 1000)
// 	fmt.Println(array)
// 	result := InsertSortAsc(array)
// 	fmt.Println(result)
// }
