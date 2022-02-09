package b_huawei_leetcode

import (
	"fmt"
	"strconv"
)

type MyStack struct {
	array []string
}

func (m *MyStack) Push(x string) {
	m.array = append(m.array, x)
}

func (m *MyStack) Pop() string {
	if len(m.array) == 0 {
		return ""
	}
	x := m.array[len(m.array)-1]
	m.array = m.array[:len(m.array)-1]
	return x
}

// 35[a2[c]]
func (m *MyStack) Empty() bool {
	return len(m.array) == 0
}

func (m *MyStack) Clear() {
	m.array = []string{}
}

func (m *MyStack) GetNum() int {
	if m.Empty() {
		return 0
	}

	ten, result := 1, 0
	for i := len(m.array) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(m.array[i])
		result += num * ten
		ten *= 10
	}
	return result
}

//3[a]
// aaabcbc
// bcaaabcaaa
func decodeString(s string) string {
	strStack := MyStack{}
	numStack := MyStack{}
	for _, r := range s {
		fmt.Println(string(r))
		fmt.Println(strStack.array)
		if IsDigit(r) {
			numStack.Push(string(r))
		}
		if IsAlphm(r) {
			strStack.Push(string(r))
		}
		if IsLeft(r) {
			if !numStack.Empty() {
				strStack.Push(strconv.Itoa(numStack.GetNum()))
				numStack.Clear()
			}
			strStack.Push(string(r))
		}
		if IsRight(r) {
			tempStr := ""
			res := ""
			for x := strStack.Pop(); x != "["; {
				fmt.Println(x)
				tempStr = x + tempStr
				x = strStack.Pop()
			}
			num, _ := strconv.Atoi(strStack.Pop())
			for i :=0; i < num; i ++ {
				res += tempStr
			}
			strStack.Push(res)
		}
	}
	var result string
	for x := strStack.Pop(); x != ""; {
		result  = x + result
		x = strStack.Pop()
	}
	return result
}

func IsDigit(x rune) bool {
	return int(x) >= 48 && int(x) <= 57
}

func IsAlphm(x rune) bool {
	i := int(x)
	return (i >= 65 && i <= 90) || (i >= 97 && i <= 122)
}

func IsRight(x rune) bool {
	return int(x) == 93
}

func IsLeft(x rune) bool {
	return int(x) == 91
}
