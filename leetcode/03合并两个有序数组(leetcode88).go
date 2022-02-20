package leetcode

import (
	"fmt"
)

/*
https://leetcode-cn.com/problems/merge-sorted-array
1. 遍历两个数组，直到一个为空
2. 创建数组长度和的新数组，全都自减一
*/

// func merge(nums1 []int, m int, nums2 []int, n int)  {
// 	for i := 0; i <n; i ++ {
// 		nums1[m + i] = nums2[i]
// 	}
// 	sort.Ints(nums1)
// }

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	var i, j int
// 	var arr []int
// 	for i < m && j < n {
// 		if nums1[i] < nums2[j] {
// 			arr = append(arr, nums1[i])
// 			i++
// 		} else {
// 			arr = append(arr, nums2[j])
// 			j++
// 		}
// 	}
// 	if i == m {
// 		arr = append(arr, nums2[j:]...)
// 	} else {
// 		arr = append(arr, nums1[i:]...)
// 	}
// 	copy(nums1, arr)
// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	var i, j = m - 1, n - 1
	for ; index >= 0; index-- {
		if i < 0 {
			nums1[index] = nums2[j]
			j--
		} else if j < 0 {
			nums1[index] = nums1[i]
			i--
		} else if nums1[i] < nums2[j] {
			nums1[index] = nums2[j]
			j--
		} else {
			nums1[index] = nums1[i]
			i--
		}
	}
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
