package main

func CountSort(array []int) []int {
	var max int
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	tempArr := make([]int, max+1)
	for i := 0; i < len(array); i++ {
		tempArr[array[i]]++
	}
	for j, d := 0, 0; j < len(tempArr); j++ {
		for k := tempArr[j]; k > 0; k-- {
			array[d] = j
			d++
		}
	}
	return array
}

// func main() {
// 	array := helper.RandomIntArray(100000, 1000)
// 	newArray := make([]int, len(array))
// 	for index, value := range array {
// 		newArray[index] = value
// 	}
// 	result := CountSort(newArray)
// 	sort.Ints(array)
// 	for index, value := range array {
// 		if value != result[index] {
// 			fmt.Println(index)
// 			panic("结果不一样")
// 		}
// 	}
// }
