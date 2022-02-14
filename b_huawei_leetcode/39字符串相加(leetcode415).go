package b_huawei_leetcode

import "strconv"

// num1 = "456", num2 = "77"
func addStrings(num1 string, num2 string) string {
	if len(num1) == 0 {
		return num2
	}
	if len(num2) == 0 {
		return num1
	}
	i, j, carry, result := len(num1)-1, len(num2)-1, 0, ""
	for i > 0 || j > 0 || carry > 0{
		var ni, nj int
		if i > 0 {
			ni = int(num1[i] - '0')
		}
		if j > 0 {
			nj = int(num2[j] - '0')
		}
		n := (ni + nj + carry) % 10
		result = strconv.Itoa(n) + result
		carry = (ni + nj + carry) / 10
		j--
		i--
	}
	return result
}
