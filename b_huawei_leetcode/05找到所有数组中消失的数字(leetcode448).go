package b_huawei_leetcode

import "fmt"

// func findDisappearedNumbers(nums []int) []int {
// 	var result []int
// 	for i := 1; i <= len(nums); i++ {
// 		var exists bool
// 		for _, value := range nums {
// 			if value == i {
// 				exists = true
// 				break
// 			}
// 		}
// 		if !exists {
// 			result = append(result, i)
// 		}
// 	}
// 	return result
// }

// func findDisappearedNumbers(nums []int) []int {
// 	var TempMap = make(map[int]struct{})
// 	for _, value := range nums {
// 		TempMap[value] = struct{}{}
// 	}
// 	var result []int
// 	for i := 1; i <= len(nums); i++ {
// 		if _, ok := TempMap[i]; !ok {
// 			result = append(result, i)
// 		}
// 	}
// 	return result
// }

func findDisappearedNumbers(nums []int) []int {
	var n = len(nums)
	for _, value := range nums {
		x := (value - 1) % n
		fmt.Println(x)
		nums[x] += n
	}
	var res []int
	for i, value := range nums {
		if value <= n {
			res = append(res, i+1)
		}
	}
	return res
}

func FindDisappearedNumbers(nums []int) {
	result := findDisappearedNumbers(nums)
	fmt.Println(result)
}
