package main

func MergeSort(array []int, left, right int) []int {
	if right == left {
		return array
	}
	mid := left + (right-left)/2
	MergeSort(array, left, mid)
	MergeSort(array, mid+1, right)
	merge(array, left, mid+1, right)
	return array
}

func merge(arr []int, left, right, bound int) {
	mid := right - 1
	temp := make([]int, bound-left+1)

	var i, j, k = left, right, 0
	// 重合长度数组合并
	for i <= mid && j <= bound {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			k++
			i++
		} else {
			temp[k] = arr[j]
			
			j++
			k++
		}
	}

	// 剩余长度数组合并
	for ; i <= mid; i++ {
		temp[k] = arr[i]
		k++
	}

	for ; j <= bound; j++ {
		temp[k] = arr[j]
		k++
	}

	// 新建数组还原到原来数组
	for index, value := range temp {
		arr[left+index] = value
	}
}

// func main() {
// 	array := helper.RandomIntArray(10, 1000)
// 	fmt.Println(array)
// 	result := MergeSort(array, 0, len(array)-1)
// 	fmt.Println(result)

// 	a := 1
// 	b := 2
// 	c := 3

// 	if a > b{
// 		if a > c{
// 			fmt.Println(a)
// 		}else {
// 			fmt.Println(c)
// 		}
// 	} else {
// 		if b < c {
// 			fmt.Println(c)
// 		} else{
// 			fmt.Println(b)
// 		}
// 	}
// }
