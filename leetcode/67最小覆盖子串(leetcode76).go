package leetcode

import (
	"fmt"
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
    ori, cnt := map[byte]int{}, map[byte]int{}
    for i := 0; i < len(t); i++ {
        ori[t[i]]++
    }
	fmt.Println(ori)

    sLen := len(s)
    len := math.MaxInt32
    ansL, ansR := -1, -1

    check := func() bool {
        for k, v := range ori {
            if cnt[k] < v {
                return false
            }
        }
        return true
    }
    for l, r := 0, 0; r < sLen; r++ {
        if r < sLen && ori[s[r]] > 0 {
            cnt[s[r]]++
        }
        for check() && l <= r {
			
            if (r - l + 1 < len) {
                len = r - l + 1
                ansL, ansR = l, l + len
            }
            if _, ok := ori[s[l]]; ok {
                cnt[s[l]] -= 1
            }
            l++
        }
    }
    if ansL == -1 {
        return ""
    }
    return s[ansL:ansR]
}

