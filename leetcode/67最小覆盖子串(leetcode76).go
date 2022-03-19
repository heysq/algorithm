package leetcode

import (
	"math"
)

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
