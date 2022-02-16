package b_huawei_leetcode

import "fmt"

/*
https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array
1. 暴力解法 循环n，循环nums，不存在记录到结果数组
2. map存储nums数组存在的元素，循环n，查map，不存在就记录到结果数组
3. 用nums本身当作hash表统计，将数组元素对应为索引的位置加n,遍历加n后的数组，若数组元素值小于等于n，则说明数组下标值不存在，即消失的数字
*/
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
	for _, v := range nums {
		v := (v - 1) % n
		fmt.Println(v)
		nums[v] += n
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
