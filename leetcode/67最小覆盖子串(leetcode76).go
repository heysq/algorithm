package leetcode

import (
	"math"
)

// func minWindow(s string, t string) string {
// 	if len(s) == 0 || len(t) == 0 || (len(s) < len(t)) {
// 		return ""
// 	}

// 	if s == t {
// 		return s
// 	}
// 	sMap := make(map[rune]int)
// 	tMap := make(map[rune]int)
// 	for _, i := range s {
// 		sMap[i] = 0
// 	}
// 	for _, i := range t {
// 		tMap[i] = 1
// 	}
// 	var left, right, min = 0, 0, math.MaxInt
// 	var result string
// 	for right < len(s) {
// 		c := rune(s[right])

// 		if _, ok := tMap[c]; ok {
// 			sMap[c] += 1
// 		}

// 		for contain(tMap, sMap) {
// 			if right-left < min {
// 				min = right - left
// 				result = s[left : right+1]
// 			}
// 			sMap[rune(s[left])] -= 1
// 			left++
// 		}
// 		right++
// 	}
// 	return result
// }

// // tMap, sMap
// func contain(src, dst map[rune]int) bool {
// 	for key, val := range src {
// 		if dst[key] < val {
// 			return false
// 		}
// 	}
// 	return true
// }

func minWindow(s string, t string) string {
	var window = make(map[byte]int)
	var need = make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	var left, right, valid, start int
	var size = math.MaxInt
	for right < len(s) {
		c := s[right]
		right++
		if val, ok := need[c]; ok {
			window[c]++
			if window[c] == val {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < size {
				start = left
				size = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if size == math.MaxInt {
		return ""
	}
	return s[start : start+size]
}
