package main


func RadixSort(array []int) []int {
	var max int
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	for i, divisor := 0, 1; max/divisor > 0; i++ {
		tempMap := make(map[int][]int)
		for _, value := range array {
			radix := (value / divisor) % 10
			if _, ok := tempMap[radix]; !ok {
				tempMap[radix] = []int{value}
			} else {
				tempMap[radix] = append(tempMap[radix], value)
			}
		}
		divisor = 10 * divisor

		for j, k := 0, 0; j < 10; j++ {
			for _, value := range tempMap[j] {
				array[k] = value
				k++
			}
		}
	}
	return array
}

// func main() {
// 	array := helper.RandomIntArray(10000, 100000)
// 	newArray := make([]int, len(array))
// 	for index, value := range array {
// 		newArray[index] = value
// 	}
// 	result := RadixSort(newArray)
// 	//fmt.Println(result)
// 	sort.Ints(array)
// 	for index, value := range array {
// 		if value != result[index] {
// 			fmt.Println(index)
// 			panic("结果不一样")
// 		}
// 	}
// }
