package leetcode

import (
	"container/list"
)

func openLock(deadends []string, target string) int {
	if contains(deadends, target) {
		return -1
	}
	queue := list.New()
	queue.PushBack("0000")
	var visited = make(map[string]struct{})
	visited["0000"] = struct{}{}
	var step int
	for size := queue.Len(); size != 0; size = queue.Len() {
		for i := 0; i < size; i++ {
			str, _ := queue.Remove(queue.Front()).(string)
			if str == target {
				return step
			}
			if contains(deadends, str) {
				continue
			}
			for j := 0; j < 4; j++ {
				upStr := up(str, j)
				if _, ok := visited[upStr]; !ok {
					queue.PushBack(upStr)
					visited[upStr] = struct{}{}
				}
				downStr := down(str, j)
				if _, ok := visited[downStr]; !ok {
					queue.PushBack(downStr)
					visited[downStr] = struct{}{}
				}
			}
		}
		step++
	}
	return -1
}

func up(s string, i int) string {
	if i >= len(s) {
		return s
	}
	t := s[i] + 1
	if s[i] == '9' {
		t = '0'
	}
	return s[:i] + string(t) + s[i+1:]
}

func down(s string, i int) string {
	if i >= len(s) {
		return s
	}
	t := s[i] - 1
	if s[i] == '0' {
		t = '9'
	}
	return s[:i] + string(t) + s[i+1:]
}

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
