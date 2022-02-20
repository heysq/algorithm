package leetcode

import (
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

/*

 */
//3[a]
// aaabcbc
// bcaaabcaaa
/*
1. 遇见字母压栈
2. 遇见数字压入数字栈
3. 遇见左括号，把数字栈的内容全部取出，然后压入字母栈，并将自己压入栈
4. 遇见右括,字母栈栈，遇到中括号，停止出栈，下一个数字循环个字符串，然后将结果再次入栈
5. 最后如果字符串遍历完后结束循环，如果栈还不为空，就把栈的内容弹出，拼到结果的前边
*/
func decodeString(s string) string {
	strStack := MyStack{}
	numStack := MyStack{}
	for _, r := range s {
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
				tempStr = x + tempStr
				x = strStack.Pop()
			}
			num, _ := strconv.Atoi(strStack.Pop())
			for i := 0; i < num; i++ {
				res += tempStr
			}
			strStack.Push(res)
		}
	}
	var result string
	for x := strStack.Pop(); x != ""; {
		result = x + result
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
